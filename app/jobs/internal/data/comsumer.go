package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
