package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"sync"
	"time"
)

type groupMsgRead struct {
	mu             sync.Mutex
	conversationId string
	push           *rws.Push
	pushCh         chan *rws.Push
	count          int // 未读消息数量

	pushTime time.Time // 最后一次推送时间
	done     chan struct{}
}

func newGroupMsgRead(push *rws.Push, pushCh chan *rws.Push) *groupMsgRead {
	g := &groupMsgRead{
		conversationId: push.ConversationId,
		push:           push,
		pushCh:         pushCh,
		count:          1,
		pushTime:       time.Now(),
		done:           make(chan struct{}),
	}

	go g.transfer()

	return g
}

// 合并推送
func (g *groupMsgRead) mergePush(push *rws.Push) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.count++

	if g.push == nil {
		g.push = push
		return
	}

	for k, v := range push.ReadRecords {
		// 消息重复直接替换
		g.push.ReadRecords[k] = v
	}
}

// IsIdle 是否空闲
func (g *groupMsgRead) IsIdle() bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.isIdle()
}

// 转发消息
func (g *groupMsgRead) transfer() {
	timer := time.NewTimer(groupMsgMergeInterval / 2)
	defer timer.Stop()

	for {
		select {
		case <-g.done:
			return
		case <-timer.C:
			g.mu.Lock()
			pushTime := g.pushTime
			val := groupMsgMergeInterval - time.Since(pushTime)
			push := g.push
			if val > 0 && g.count < groupMsgMergeMaxSize || push == nil {
				if val > 0 {
					timer.Reset(val)
				}
				g.mu.Unlock()
				continue
			}

			// 推送消息
			g.pushTime = time.Now()
			g.push = nil
			g.count = 0
			timer.Reset(groupMsgMergeInterval / 2)
			g.mu.Unlock()
			g.pushCh <- push
		default:
			g.mu.Lock()
			if g.count >= groupMsgMergeMaxSize {
				push := g.push
				g.pushTime = time.Now()
				g.push = nil
				g.count = 0
				g.mu.Unlock()
				g.pushCh <- push
				continue
			}

			if g.isIdle() {
				g.mu.Unlock()
				log.Info("groupMsgRead transfer idle")
				// 发送信号让consumerUsecase清理数据
				g.pushCh <- &rws.Push{
					ChatType:       constants.ChatTypeGroup,
					ConversationId: g.conversationId,
				}
				continue
			}
			// 活跃中
			g.mu.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (g *groupMsgRead) isIdle() bool {
	pushTime := g.pushTime
	val := groupMsgMergeInterval*2 - time.Since(pushTime)
	if val <= 0 && g.push == nil && g.count == 0 {
		// 空闲
		return true
	}

	return false
}

func (g *groupMsgRead) clear() {
	select {
	case <-g.done:
	default:
		close(g.done)
	}

	g.push = nil
}
