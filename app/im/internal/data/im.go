package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kratos-im/app/im/internal/biz"
	"kratos-im/model"

	"github.com/go-kratos/kratos/v2/log"
)

type imRepo struct {
	data *Data
	log  *log.Helper
}

// NewIMRepo .
func NewIMRepo(data *Data, logger log.Logger) biz.IMRepo {
	return &imRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *imRepo) Save(ctx context.Context, chatLog model.ChatLog) error {
	_, err := r.data.mongoDatabase.Collection(chatLog.Collection()).InsertOne(ctx, chatLog)
	if err != nil {
		return err
	}
	return nil
}

func (r *imRepo) ListBySendTime(ctx context.Context, conversationId string, startSendTime, endSendTime, limit int64) ([]*model.ChatLog, error) {
	var data []*model.ChatLog

	opt := options.FindOptions{
		Limit: &biz.DefaultChatLogCount,
		Sort: bson.M{
			"sendTime": -1,
		},
	}

	filter := bson.M{
		"conversationId": conversationId,
	}
	if endSendTime > 0 {
		//  startSendTime > x endSendTime
		filter["sendTime"] = bson.M{
			"$gt":  endSendTime,
			"$lte": startSendTime,
		}
	} else {
		filter["sendTime"] = bson.M{
			"$lt": startSendTime,
		}
	}

	if limit > 0 {
		opt.Limit = &limit
	}

	cur, err := r.data.mongoDatabase.Collection(model.ChatLog{}.Collection()).Find(ctx, filter, &opt)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *imRepo) FindChatLogOne(ctx context.Context, id primitive.ObjectID) (*model.ChatLog, error) {
	var data model.ChatLog

	err := r.data.mongoDatabase.Collection(model.ChatLog{}.Collection()).
		FindOne(ctx, bson.D{{"_id", id}}).
		Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *imRepo) FindConversationOne(ctx context.Context, id string) (*model.Conversation, error) {
	var data model.Conversation

	err := r.data.mongoDatabase.Collection(model.Conversation{}.Collection()).
		FindOne(ctx, bson.M{"conversationId": id}).
		Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *imRepo) ListByConversationIds(ctx context.Context, ids []string) ([]*model.Conversation, error) {
	var data []*model.Conversation

	cur, err := r.data.mongoDatabase.Collection(model.Conversation{}.Collection()).Find(ctx, bson.M{
		"conversationId": bson.M{
			"$in": ids,
		},
	})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *imRepo) ConversationsByUserId(ctx context.Context, uid string) (*model.Conversations, error) {
	var data model.Conversations

	err := r.data.mongoDatabase.Collection(model.Conversations{}.Collection()).FindOne(ctx, bson.M{
		"userId": uid,
	}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *imRepo) UpdateMsg(ctx context.Context, chatLog *model.ChatLog) error {
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

func (r *imRepo) UpdateConversations(ctx context.Context, data *model.Conversations) error {
	_, err := r.data.mongoDatabase.Collection(model.Conversations{}.Collection()).UpdateOne(ctx,
		bson.D{{"_id", data.ID}},
		bson.M{
			"$set": data,
		},
		options.Update().SetUpsert(true), // 如果不存在则插入
	)
	return err
}

func (r *imRepo) CreateConversation(ctx context.Context, conversation model.Conversation) error {
	_, err := r.data.mongoDatabase.Collection(model.Conversation{}.Collection()).InsertOne(context.Background(), conversation)
	if err != nil {
		return err
	}
	return nil
}
