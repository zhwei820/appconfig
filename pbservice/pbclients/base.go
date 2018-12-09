package pbclients

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"sync"
	"sync/atomic"
)

type ClientRpcs struct {
	Idx       int64
	DownIdx   int64
	DownPos   int64
	Mutex     sync.RWMutex
	Clients   []*rpc.HTTPClient
	Services  []interface{}
	skipTimes int64
}

func CreateNewInstance() *ClientRpcs {
	skipTimes := beego.AppConfig.DefaultInt64("rpc_skip_times", 50)
	return &ClientRpcs{
		Idx:       0,
		DownIdx:   -1,
		DownPos:   -1,
		skipTimes: skipTimes,
		Mutex:     sync.RWMutex{},
		Clients:   make([]*rpc.HTTPClient, 0),
		Services:  make([]interface{}, 0),
	}
}

func (c *ClientRpcs) GetService() (svc interface{}, idx, pos int64, err error) {
	// get svc instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.Services) > 0 {
		c.switchService()
		pos := atomic.LoadInt64(&(c.Idx)) % int64(len(c.Services))
		if pos == atomic.LoadInt64(&(c.DownIdx)) {
			c.switchService()
			pos = atomic.LoadInt64(&(c.Idx)) % int64(len(c.Services))
		}
		if c.Idx > c.DownIdx+int64(c.skipTimes) { // 尝试将这个节点恢复
			c.BringUpService()
		}
		return c.Services[pos], c.Idx, pos, nil
	}

	return nil, 0, 0, errors.New("empty svc")
}

func (c *ClientRpcs) switchService() {
	atomic.AddInt64(&c.Idx, 1)
}

func (c *ClientRpcs) KnockdownService(idx, pos int64) {
	// 暂时将这个节点移除
	atomic.StoreInt64(&c.DownIdx, idx)
	atomic.StoreInt64(&c.DownPos, pos)
}

func (c *ClientRpcs) BringUpService() {
	// 将这个节点恢复
	atomic.StoreInt64(&c.DownIdx, -1)
	atomic.StoreInt64(&c.DownPos, -1)
}
