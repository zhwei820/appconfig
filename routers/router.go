package routers

import (
	_ "github.com/zhwei820/appconfig/utils/util"
	_ "github.com/zhwei820/appconfig/pbservice/services"

	"github.com/astaxie/beego"
	"github.com/zhwei820/appconfig/services/default_service"
	"github.com/zhwei820/appconfig/services/group_service"
	_ "github.com/astaxie/beego/session/redis"

	"github.com/zhwei820/appconfig/services/staffuser_service"
	"github.com/zhwei820/appconfig/services/auth_service"
)

// @APIVersion 1.0.0
// @Description beego Test API
// 使用注释路由
// @SecurityDefinition jwt apiKey Authorization header header
func init() {
	// 基础初始化
	BaseInit()

	beego.Router("/", &default_service.DefaultController{}, "*:ApiGetAll")
	beego.Router("/register", &auth_service.AuthController{}, "post:ApiRegister")
	beego.Router("/api_login", &auth_service.AuthController{}, "post:ApiLogin")
	beego.Router("/login", &auth_service.AuthController{}, "post:SessionLogin")
	beego.Router("/logout", &auth_service.AuthController{}, "post:SessionLogout")

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
