package pbservice

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
)

func init() {

	service := rpc.NewHTTPService()
	service.AddFunction("hello", Say)
	beego.Handler("/rpc", service)
}
