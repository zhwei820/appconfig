package services

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/pbservice/pb/appconfig/say"
	. "github.com/zhwei820/appconfig/pbservice/pb/appconfig/sing"
	"github.com/zhwei820/appconfig/pbservice/services/sing"
	"github.com/zhwei820/appconfig/pbservice/services/say"
)

func init() {

	service := rpc.NewHTTPService()
	service.AddInstanceMethods(SayService{SayHello: say.SayHello}, rpc.Options{Simple: true})
	service.AddInstanceMethods(SingService{SingHello: sing.SingHello}, rpc.Options{Simple: true})
	beego.Handler(define.DiscoveryUrlPrefix, service)
}
