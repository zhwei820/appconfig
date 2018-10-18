package pbservice

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/pb/appconfig/sing"
	. "github.com/zhwei820/appconfig/pb/appconfig/say"
	"github.com/zhwei820/appconfig/pbservice/sing"
	"github.com/zhwei820/appconfig/pbservice/say"
)

func init() {

	service := rpc.NewHTTPService()
	service.AddInstanceMethods(SayService{Hello: say.SayHello}, rpc.Options{Simple: true})
	service.AddInstanceMethods(SingService{Hello: sing.SingHello}, rpc.Options{Simple: true})
	beego.Handler(define.DiscoveryUrlPrefix, service)
}
