package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"github.com/hprose/hprose-golang/rpc"
	"strings"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	"time"
)

type consumer struct {
	idx         int
	conf        *naming.Config
	appID       string
	Dis         naming.Resolver
	ins         []*naming.Instance
	rpclients   []*rpc.HTTPClient
	rpcservices []interface{}
}

// This Example show how get watch a server provier and get provider instances.
func NewComsumer(appID string) *consumer {
	discoveryUrls := strings.Split(beego.AppConfig.String("discovery_url"), ",")

	conf := &naming.Config{
		Nodes: discoveryUrls,
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
	}
	dis := naming.New(conf)
	demoComsumer := &consumer{
		conf:  conf,
		appID: appID,
		Dis:   dis.Build(appID),
	}
	go renewInstances(demoComsumer)
	return demoComsumer
}

// 定时更新服务列表
func renewInstances(demoComsumer *consumer) {
	renewfunc := func() {
		ch := demoComsumer.Dis.Watch()
		demoComsumer.getInstances(ch)
	}

	go renewfunc()
	timer2 := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-timer2.C:
			renewfunc()
		}
	}

}
