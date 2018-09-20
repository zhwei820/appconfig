package pbclients

import (
	"fmt"
	"time"

	"github.com/Bilibili/discovery/naming"
)


type consumer struct {
	conf  *naming.Config
	appID string
	dis   naming.Resolver
	ins   []*naming.Instance
}

// This Example show how get watch a server provier and get provider instances.
func ExampleDiscovery_Watch() {
	conf := &naming.Config{
		Nodes: []string{"127.0.0.1:7171"},
		Zone:  "sh1",
		Env:   "test",
	}
	dis := naming.New(conf)
	c := &consumer{
		conf:  conf,
		appID: "provider",
		dis:   dis.Build("provider"),
	}
	rsl := dis.Build(c.appID)
	ch := rsl.Watch()
	go c.getInstances(ch)
	in := c.getInstance()
	_ = in
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

func (c *consumer) getInstance() (ins *naming.Instance) {
	// get instance by loadbalance
	// you can use any loadbalance algorithm what you want.
	return
}
