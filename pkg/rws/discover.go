package rws

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var _ Discover = (*nopDiscover)(nil)

type Discover interface {
	// Register 注册服务
	Register(serverAddr string) error
	// BoundUser 绑定用户
	BoundUser(uid string) error
	// RelieveUser 解除与用户绑定
	RelieveUser(uid string) error
	// Transpond 转发
	Transpond(msg interface{}, uid ...string) error
}

// 默认的
type nopDiscover struct {
	serverAddr string
}

func NewNopDiscover() Discover {
	return &nopDiscover{}
}

// 注册服务
func (d *nopDiscover) Register(serverAddr string) error { return nil }

// 绑定用户
func (d *nopDiscover) BoundUser(uid string) error { return nil }

func (d *nopDiscover) RelieveUser(uid string) error { return nil }

// 转发消息
func (d *nopDiscover) Transpond(msg interface{}, uid ...string) error { return nil }

// 默认的
type redisDiscover struct {
	serverAddr   string
	patten       string
	auth         http.Header
	srvKey       string
	boundUserKey string
	redis        redis.Cmdable
	clients      map[string]IClient
}

func NewRedisDiscover(auth http.Header, srvKey string, cfg *redis.Options) Discover {
	return &redisDiscover{
		patten:       "/ws",
		srvKey:       srvKey,
		boundUserKey: fmt.Sprintf("%s.%s", srvKey, "boundUserKey"),
		redis:        redis.NewClient(cfg),
		clients:      make(map[string]IClient),
		auth:         auth,
	}
}

// 注册服务
func (d *redisDiscover) Register(serverAddr string) (err error) {
	d.serverAddr = serverAddr
	ctx := context.Background()

	// 服务列表：redis存储用set
	go d.redis.SAdd(ctx, d.srvKey, serverAddr)

	return
}

// 绑定用户
func (d *redisDiscover) BoundUser(uid string) (err error) {
	// 用户绑定
	ctx := context.Background()
	//exists, err := d.redis.HExists(ctx, d.boundUserKey, uid).Result()
	//if err != nil {
	//	return err
	//}
	//if exists {
	//	// 存在绑定关系
	//	return nil
	//}

	// 绑定
	return d.redis.HSet(ctx, d.boundUserKey, uid, d.serverAddr).Err()
}

func (d *redisDiscover) RelieveUser(uid string) (err error) {
	ctx := context.Background()
	_, err = d.redis.HDel(ctx, d.boundUserKey, uid).Result()
	return
}

// 转发消息
func (d *redisDiscover) Transpond(msg interface{}, uids ...string) (err error) {
	ctx := context.Background()
	for _, uid := range uids {
		srvAddr, err := d.redis.HGet(ctx, d.boundUserKey, uid).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
		if srvAddr == "" {
			continue
		}
		srvClient, ok := d.clients[srvAddr]
		if !ok {
			srvClient, err = d.createClient(srvAddr, d.patten)
			if err != nil {
				log.Errorf("create client error: %v", err)
				return err
			}
			d.clients[srvAddr] = srvClient
		}

		log.Info("redis transpand -》 ", srvAddr, " uid ", uid)

		if err = d.send(srvClient, msg, uid); err != nil {
			return err
		}
	}

	d.clients = make(map[string]IClient)

	return
}

func (d *redisDiscover) send(srvClient IClient, msg interface{}, uid string) error {
	return srvClient.Send(Message{
		FrameType:    FrameTranspond,
		TranspondUid: uid,
		Data:         msg,
	})
}

func (d *redisDiscover) createClient(srvAddr string, patten string) (IClient, error) {
	return NewClient(srvAddr, patten, d.auth)
}
