package pbservice

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/utils/define"
)

func init() {

	service := rpc.NewHTTPService()
	service.AddFunction("hello", Say)
	beego.Handler(define.DiscoveryUrlPrefix, service)
}
