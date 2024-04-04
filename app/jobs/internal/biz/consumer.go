package biz

import (
	"context"
	"encoding/base64"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kratos-im/app/jobs/internal/conf"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/rws"
	"kratos-im/pkg/tools"
	"sync"
	"time"
)

const (
	defaultGroupMsgMerge = false // 默认不合并

)

var (
	groupMsgMergeInterval = 10 * time.Second // 默认10s
	groupMsgMergeMaxSize  = 10
)

// ConsumerRepo is a Greater repo.
type ConsumerRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
	UpdateMsg(ctx context.Context, chatLog *model.ChatLog) error
	ListGroupMembersByGid(ctx context.Context, gid uint64) ([]*model.GroupMembers, error)
	ListChatLogByIds(ctx context.Context, msgids []string) ([]*model.ChatLog, error)
	UpdateRead(ctx context.Context, id primitive.ObjectID, readRecords []byte) error
}

// ConsumerUsecase is a Consumer usecase.
type ConsumerUsecase struct {
	repo            ConsumerRepo
	log             *log.Helper
	baseMsgTransfer *BaseMsgTransfer
	msgReadHandle   *conf.MsgReadHandler
	mu              sync.Mutex
	groupMsgReads   map[string]*groupMsgRead
	pushCh          chan *rws.Push
}

// NewConsumerUsecase new a Consumer usecase.
func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger, baseMsgTransfer *BaseMsgTransfer, msgReadHandle *conf.MsgReadHandler) *ConsumerUsecase {
	if msgReadHandle == nil {
		msgReadHandle = &conf.MsgReadHandler{GroupMsgMerge: defaultGroupMsgMerge}
	}
	if msgReadHandle.GroupMsgMerge {
		if msgReadHandle.GroupMsgMergeInterval > 0 {
			groupMsgMergeInterval = time.Duration(msgReadHandle.GroupMsgMergeInterval) * time.Second
		}
		if msgReadHandle.GroupMsgMergeMaxSize > 0 {
			groupMsgMergeMaxSize = int(msgReadHandle.GroupMsgMergeMaxSize)
		}
	}

	uc := &ConsumerUsecase{
		repo:            repo,
		log:             log.NewHelper(logger),
		baseMsgTransfer: baseMsgTransfer,
		msgReadHandle:   msgReadHandle,
		groupMsgReads:   make(map[string]*groupMsgRead, 1),
		pushCh:          make(chan *rws.Push, 1),
	}

	go uc.transfer()

	return uc
}

// HandleMsgTransfer 转发消息
func (u *ConsumerUsecase) HandleMsgTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgChatTransfer) error {
	chatLog := model.ChatLog{
		ID:             primitive.NewObjectID(),
		ConversationId: msg.ConversationId,
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		MsgFrom:        0,
		ChatType:       msg.ChatType,
		MsgType:        msg.MType,
		MsgContent:     msg.Content,
		SendTime:       msg.SendTime,
		Status:         0,
	}

	// 自身已读
	readRecords := tools.NewBitmap(0)
	readRecords.Set(chatLog.SendId)
	chatLog.ReadRecords = readRecords.Export()

	// 保存数据
	err := u.repo.Save(ctx, chatLog)
	if err != nil {
		u.log.Errorf("HandleMsgTransfer Save: %v", err)
		return err
	}

	// 更新会话
	err = u.repo.UpdateMsg(ctx, &chatLog)
	if err != nil {
		u.log.Errorf("HandleMsgTransfer UpdateMsg: %v", err)
		return err
	}

	// 推送消息(推送给 ws server)
	return u.baseMsgTransfer.Transfer(ctx, &rws.Push{
		MsgId:          chatLog.ID.Hex(),
		ConversationId: msg.ConversationId,
		ChatType:       msg.ChatType,
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		RecvIds:        msg.RecvIds,
		MType:          msg.MType,
		Content:        msg.Content,
		SendTime:       msg.SendTime,
	})
}

// HandleMsgReadTransfer 已读消息
func (u *ConsumerUsecase) HandleMsgReadTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgMarkReadTransfer) error {
	// 查询聊天记录
	chatLogs, err := u.repo.ListChatLogByIds(ctx, msg.MsgIds)
	if err != nil {
		return err
	}

	// 已读记录
	res := make(map[string]string)
	for _, v := range chatLogs {
		switch v.ChatType {
		case constants.ChatTypeSingle: // 单聊
			v.ReadRecords = []byte{1}
		case constants.ChatTypeGroup: // 群聊
			// 更新已读记录
			readRecords := tools.Load(v.ReadRecords)
			readRecords.Set(msg.SendId) // 设置已读
			v.ReadRecords = readRecords.Export()
		}

		// 转string保证精度
		res[v.ID.Hex()] = base64.StdEncoding.EncodeToString(v.ReadRecords)

		// 更新已读
		err = u.repo.UpdateRead(ctx, v.ID, v.ReadRecords)
		if err != nil {
			return err
		}
	}
	u.log.Infof("已读记录: %s, headers: %v, msg: %v", topic, headers, msg)

	push := &rws.Push{
		ConversationId: msg.ConversationId,
		ChatType:       msg.ChatType,
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		ContentType:    constants.ContentTypeMakeRead,
		ReadRecords:    res,
	}

	switch push.ChatType {
	case constants.ChatTypeSingle: // 私聊(直接推送)
		u.pushCh <- push
	case constants.ChatTypeGroup: // 群聊
		// 是否开启合并推送
		if u.msgReadHandle.GroupMsgMerge == defaultGroupMsgMerge {
			// 不合并推送
			u.pushCh <- push
		}

		u.mu.Lock()
		defer u.mu.Unlock()

		push.SendId = ""
		if _, ok := u.groupMsgReads[push.ConversationId]; !ok {
			u.groupMsgReads[push.ConversationId] = newGroupMsgRead(push, u.pushCh)
		} else {
			// 合并推送
			u.groupMsgReads[push.ConversationId].mergePush(push)
			u.log.Infof("合并推送: %s, headers: %v, msg: %v", topic, headers, msg)
		}
	}

	// 推送消息(推送给 ws server)
	return nil
}

func (u *ConsumerUsecase) transfer() {
	for p := range u.pushCh {
		if p.RecvId != "" || len(p.RecvIds) > 0 {
			err := u.baseMsgTransfer.Transfer(context.Background(), p)
			if err != nil {
				u.log.Errorf("transfer: %v", err)
			}
		}

		if p.ChatType == constants.ChatTypeSingle {
			// 单聊
			continue
		}

		if u.msgReadHandle.GroupMsgMerge == defaultGroupMsgMerge {
			// 不合并推送
			continue
		}

		// 清空数据
		u.mu.Lock()
		if _, ok := u.groupMsgReads[p.ConversationId]; ok && u.groupMsgReads[p.ConversationId].IsIdle() {
			// 空闲
			u.groupMsgReads[p.ConversationId].clear()
			delete(u.groupMsgReads, p.ConversationId)
		}
		u.mu.Unlock()
	}
}
