package pbclients

import (
	"github.com/Bilibili/discovery/naming"
	"github.com/hprose/hprose-golang/rpc"
	"strings"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	"sync"
)

type Consumer struct {
	Idx   uint64
	mutex sync.RWMutex

	conf        *naming.Config
	appID       string
	Dis         naming.Resolver
	ins         []*naming.Instance
	rpclients   []*rpc.HTTPClient
	rpcservices []interface{}
}

// This Example show how get watch a server provier and get provider instances.
func NewComsumer(appID string) *Consumer {
	discoveryUrls := strings.Split(beego.AppConfig.String("discovery_url"), ",")

	conf := &naming.Config{
		Nodes: discoveryUrls,
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
	}
	dis := naming.New(conf)
	demoComsumer := &Consumer{
		conf:  conf,
		appID: appID,
		Dis:   dis.Build(appID),
	}
	ch := demoComsumer.Dis.Watch()
	go demoComsumer.getInstances(ch)
	return demoComsumer
}
