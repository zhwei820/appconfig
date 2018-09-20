package pbservice

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
)

func init() {
	println("init1")

	service := rpc.NewHTTPService()
	service.AddFunction("hello", Say)
	beego.Handler("/rpc", service)
}
