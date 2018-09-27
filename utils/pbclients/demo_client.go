package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/pb/appconfig"
	"github.com/hprose/hprose-golang/rpc"
	"errors"
	"sync/atomic"
)

var DemoComsumer *Consumer

func init() {
	DemoComsumer = NewComsumer(define.SingServiceAppId)
}

func (c *Consumer) getInstances(ch <-chan struct{}) {

	for { // NOTE: 通过watch返回的event chan =>
		if _, ok := <-ch; !ok {
			return
		}
		// NOTE: <= 实时fetch最新的instance实例
		ins, ok := c.Dis.Fetch()
		if !ok {
			continue
		}
		// get local zone instances
		if in, ok := ins[c.conf.Zone]; ok {
			c.ins = in
			c.rpclients = make([]*rpc.HTTPClient, 0)
			c.rpcservices = make([] interface{}, 0)

			c.mutex.Lock()
			for _, item := range c.ins {
				rpclient := rpc.NewHTTPClient(item.Addrs[0] + define.DiscoveryUrlPrefix)
				c.rpclients = append(c.rpclients, rpclient)

				var singService *SingService
				rpclient.UseService(&singService)
				c.rpcservices = append(c.rpcservices, singService)
			}
			c.mutex.Unlock()

		}
	}
}

func (c *Consumer) GetInstance() (in *naming.Instance, err error) {
	// get instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.ins) > 0 {
		return c.ins[0], nil
	}
	return nil, errors.New("empty client")
}

func (c *Consumer) GetService() (svc interface{}, err error) {
	// get svc instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.rpcservices) > 0 {
		c.switchService()
		return c.rpcservices[atomic.LoadUint64(&(c.Idx))%uint64(len(c.rpcservices))], nil
	}
	return nil, errors.New("empty svc")
}
func (c *Consumer) switchService() {
	atomic.AddUint64(&c.Idx, 1)
}

func (c *Consumer) RemoveService(idx uint64) {
	c.mutex.Lock()
	c.rpcservices = append(c.rpcservices[:idx%uint64(len(c.rpcservices))], c.rpcservices[idx%uint64(len(c.rpcservices))+1:]...)
	c.mutex.Unlock()
}
