package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kratos-im/api/social"
	"kratos-im/model"

	"kratos-im/app/jobs/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type consumerRepo struct {
	data *Data
	log  *log.Helper
}

// NewConsumerRepo .
func NewConsumerRepo(data *Data, logger log.Logger) biz.ConsumerRepo {
	return &consumerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *consumerRepo) Save(ctx context.Context, chatLog model.ChatLog) error {
	_, err := r.data.mongoDatabase.Collection(chatLog.Collection()).InsertOne(ctx, chatLog)
	if err != nil {
		return err
	}
	return nil
}

func (r *consumerRepo) UpdateMsg(ctx context.Context, chatLog *model.ChatLog) error {
	_, err := r.data.mongoDatabase.Collection(model.Conversation{}.Collection()).UpdateOne(ctx,
		bson.M{"conversationId": chatLog.ConversationId},
		bson.M{
			// 更新会话总消息数
			"$inc": bson.M{"total": 1},
			"$set": bson.M{"msg": chatLog},
		},
	)
	return err
}

func (r *consumerRepo) ListGroupMembersByGid(ctx context.Context, gid uint64) ([]*model.GroupMembers, error) {
	resp, err := r.data.socialClient.GroupUsers(ctx, &social.GroupUsersReq{GroupId: gid})
	if err != nil {
		return nil, err
	}

	var data = make([]*model.GroupMembers, 0, len(resp.List))

	for _, user := range resp.List {
		data = append(data, &model.GroupMembers{
			GroupId: gid,
			UserId:  user.UserId,
		})
	}

	return data, nil
}

// ListChatLogByIds 根据ids查询聊天记录
func (r *consumerRepo) ListChatLogByIds(ctx context.Context, msgids []string) ([]*model.ChatLog, error) {
	var data = make([]*model.ChatLog, 0, len(msgids))
	ids := make([]primitive.ObjectID, 0, len(msgids))
	for _, msgid := range msgids {
		oid, err := primitive.ObjectIDFromHex(msgid)
		if err != nil {
			return nil, err
		}
		ids = append(ids, oid)
	}

	cur, err := r.data.mongoDatabase.Collection(model.ChatLog{}.Collection()).Find(ctx,
		bson.M{"_id": bson.M{
			"$in": ids,
		}})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateRead 更新已读
func (r *consumerRepo) UpdateRead(ctx context.Context, id primitive.ObjectID, readRecords []byte) error {
	_, err := r.data.mongoDatabase.Collection(model.ChatLog{}.Collection()).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"readRecords": readRecords}})
	if err != nil {
		return err
	}

	return nil
}
