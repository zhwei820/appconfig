package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/pb/appconfig"
	"github.com/hprose/hprose-golang/rpc"
	"errors"
)

var DemoComsumer *consumer

func init() {
	DemoComsumer = NewComsumer(define.SingServiceAppId)
}

func (c *consumer) getInstances(ch <-chan struct{}) {

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

			for _, item := range c.ins {
				rpclient := rpc.NewHTTPClient(item.Addrs[0] + define.DiscoveryUrlPrefix)
				c.rpclients = append(c.rpclients, rpclient)

				var singService *SingService
				rpclient.UseService(&singService)
				c.rpcservices = append(c.rpcservices, singService)

			}
		}
	}
}

func (c *consumer) GetInstance() (in *naming.Instance, err error) {
	// get instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.ins) > 0 {
		return c.ins[0], nil
	}
	return nil, errors.New("empty client")
}

func (c *consumer) GetService() (svc interface{}, err error) {
	// get svc instance by load balance
	// you can use any load balance algorithm what you want.

	if len(c.rpcservices) > 0 {
		c.idx += 1
		return c.rpcservices[c.idx%len(c.rpcservices)], nil
	}
	return nil, errors.New("empty svc")
}
