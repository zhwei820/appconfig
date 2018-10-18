package sing

import (
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/utils/pbclients"
	. "github.com/zhwei820/appconfig/pb/appconfig/sing"
	"github.com/hprose/hprose-golang/rpc"
	"strings"
	"github.com/astaxie/beego"
	"github.com/Bilibili/discovery/naming"
)

type SingComsumerStruct struct {
	Consumer
}

var SingComsumer *SingComsumerStruct

func init() {
	SingComsumer = NewComsumer(define.SingServiceAppId)
}

// This Example show how get watch a server provier and get provider instances.
func NewComsumer(appID string) *SingComsumerStruct {
	discoveryUrls := strings.Split(beego.AppConfig.String("discovery_url"), ",")

	conf := &naming.Config{
		Nodes: discoveryUrls,
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
	}
	dis := naming.New(conf)
	demoComsumer := &SingComsumerStruct{Consumer{
		Conf:  conf,
		AppID: appID,
		Dis:   dis.Build(appID),
	}}
	ch := demoComsumer.Dis.Watch()
	go demoComsumer.getInstances(ch)
	return demoComsumer
}

func (c *SingComsumerStruct) getInstances(ch <-chan struct{}) {
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
		if in, ok := ins[c.Conf.Zone]; ok {
			c.Ins = in
			c.Rpclients = make([]*rpc.HTTPClient, 0)
			c.Rpcservices = make([] interface{}, 0)

			c.Mutex.Lock()
			for _, item := range c.Ins {
				rpclient := rpc.NewHTTPClient(item.Addrs[0] + define.DiscoveryUrlPrefix)
				c.Rpclients = append(c.Rpclients, rpclient)

				var singService *SingService // replace this
				rpclient.UseService(&singService)
				c.Rpcservices = append(c.Rpcservices, singService)
			}
			c.Mutex.Unlock()

		}
	}
}
