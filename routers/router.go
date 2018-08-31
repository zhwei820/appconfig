package routers

import (
	_ "back/appconfig/utils/util"

	"github.com/astaxie/beego"
	"back/appconfig/services/default_service"
	"back/appconfig/services/group_service"
	_ "github.com/astaxie/beego/session/redis"

	"back/appconfig/services/staffuser_service"
	"back/appconfig/services/auth_service"
)

// @APIVersion 1.0.0
// @Description beego Test API
// 使用注释路由
// @SecurityDefinition jwt apiKey Authorization header header
func init() {
	// 基础初始化
	BaseInit()

	beego.Router("/", &default_service.DefaultController{}, "*:ApiGetAll")
	beego.Router("/register", &auth_service.AuthController{}, "POST:ApiRegister")
	beego.Router("/api_login", &auth_service.AuthController{}, "POST:ApiLogin")
	beego.Router("/login", &auth_service.AuthController{}, "POST:SessionLogin")

	beego.Router("/view/:dir([\\w]+)/:html([\\w]+).html", &default_service.DefaultController{}, "*:Html")
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/staffuser",
			beego.NSInclude(
				&staffuser_service.StaffUserController{},
			),
		),
		beego.NSNamespace("/group",
			beego.NSInclude(
				&group_service.GroupController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
