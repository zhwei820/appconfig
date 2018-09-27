package pbservice

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
	. "github.com/zhwei820/appconfig/pb/appconfig"
)

func init() {

	service := rpc.NewHTTPService()
	service.AddInstanceMethods(SingService{Hello: Hello}, rpc.Options{Simple: true})
	beego.Handler(define.DiscoveryUrlPrefix, service)
}
