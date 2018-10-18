package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"github.com/hprose/hprose-golang/rpc"
	"sync"
	"sync/atomic"
	"errors"
)

type Consumer struct {
	Idx   uint64
	Mutex sync.RWMutex

	Conf        *naming.Config
	AppID       string
	Dis         naming.Resolver
	Ins         []*naming.Instance
	Rpclients   []*rpc.HTTPClient
	Rpcservices []interface{}
}

func (c *Consumer) GetInstance() (in *naming.Instance, err error) {
	// get instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.Ins) > 0 {
		return c.Ins[0], nil
	}
	return nil, errors.New("empty client")
}

func (c *Consumer) GetService() (svc interface{}, err error) {
	// get svc instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.Rpcservices) > 0 {
		c.switchService()
		return c.Rpcservices[atomic.LoadUint64(&(c.Idx))%uint64(len(c.Rpcservices))], nil
	}
	return nil, errors.New("empty svc")
}
func (c *Consumer) switchService() {
	atomic.AddUint64(&c.Idx, 1)
}

func (c *Consumer) RemoveService(idx uint64) {
	c.Mutex.Lock()
	c.Rpcservices = append(c.Rpcservices[:idx%uint64(len(c.Rpcservices))], c.Rpcservices[idx%uint64(len(c.Rpcservices))+1:]...)
	c.Mutex.Unlock()
}
