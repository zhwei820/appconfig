package rpc_register

import (
	"github.com/Bilibili/discovery/naming"
	"time"
	"github.com/astaxie/beego"
	"strings"
	"context"
	"github.com/zhwei820/appconfig/utils/util"
	"github.com/rs/zerolog/log"
	"github.com/zhwei820/appconfig/utils/define"
)

var Cancel context.CancelFunc

// This Example register a server provider into discovery.
func DiscoveryRegister() {
	discoveryUrls := strings.Split(beego.AppConfig.String("discovery_url"), ",")

	conf := &naming.Config{
		Nodes: discoveryUrls, // NOTE: 配置种子节点(1个或多个)，client内部可根据/discovery/nodes节点获取全部node(方便后面增减节点)
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
	}
	dis := naming.New(conf)
	svcAddr := "http://" + util.GetOutboundIP() + ":" + beego.AppConfig.String("httpport")

	ins := &naming.Instance{
		Zone:  define.DiscoveryZone,
		Env:   define.DiscoveryEnv,
		AppID: define.DiscoveryAppID,
		// Hostname:"", // NOTE: hostname 不需要，会优先使用discovery new时Config配置的值，如没有则从os.Hostname方法获取！！！
		Addrs:    []string{svcAddr},
		Color:    "red",
		LastTs:   time.Now().Unix(),
		Metadata: map[string]string{"weight": "10"},
	}
	Cancel, _ = dis.Register(ins)
	// NOTE: 注意一般在进程退出的时候执行，会调用discovery的cancel接口，使实例从discovery移除
	log.Info().Msg("register")

}
