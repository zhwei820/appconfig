package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"strings"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	"github.com/hprose/hprose-golang/rpc"
)

type consumer struct {
	conf  *naming.Config
	appID string
	dis   naming.Resolver
	ins   []*naming.Instance
	Clients []*rpc.Clients
}

var DemoComsumer *consumer

// This Example show how get watch a server provier and get provider instances.
func init() {
	discoveryUrls := strings.Split(beego.AppConfig.String("discovery_url"), ",")

	conf := &naming.Config{
		Nodes: discoveryUrls,
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
	}
	dis := naming.New(conf)
	DemoComsumer = &consumer{
		conf:  conf,
		appID: define.DiscoveryAppID,
		dis:   dis.Build(define.DiscoveryAppID),
	}
	rsl := dis.Build(DemoComsumer.appID)
	ch := rsl.Watch()
	go DemoComsumer.getInstances(ch)
	//in := DemoComsumer.getInstance()
	//_ = in
}

func (c *consumer) getInstances(ch <-chan struct{}) {
	for { // NOTE: 通过watch返回的event chan =>
		if _, ok := <-ch; !ok {
			return
		}
		// NOTE: <= 实时fetch最新的instance实例
		ins, ok := c.dis.Fetch()
		if !ok {
			continue
		}
		// get local zone instances, otherwise get all zone instances.
		if in, ok := ins[c.conf.Zone]; ok {
			c.ins = in
		} else {
			for _, in := range ins {
				c.ins = append(c.ins, in...)
			}
		}
	}
}

func (c *consumer) GetClient() (ins *naming.Instance) {
	// get instance by loadbalance
	// you can use any loadbalance algorithm what you want.

	return
}
