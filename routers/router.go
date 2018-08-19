package routers

import (
	"github.com/astaxie/beego"
	"back/appconfig/services/user_service"
	"back/appconfig/services/default_service"
)

// @APIVersion 1.0.0
// @Description beego Test API
// 使用注释路由
// @SecurityDefinition jwt apiKey Authorization header header
func init() {
	beego.Router("/", &default_service.DefaultController{}, "*:ApiGetAll")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&default_service.TestController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user_service.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
