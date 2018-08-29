package routers

import (
	_ "back/appconfig/utils/util"

	"github.com/astaxie/beego"
	"back/appconfig/services/default_service"
	"back/appconfig/services/group_service"
	_ "github.com/astaxie/beego/session/redis"

	"back/appconfig/models"
	"github.com/astaxie/beego/context"
	"back/appconfig/utils/sentry"
	"back/appconfig/services/staffuser_service"
	"github.com/satori/go.uuid"
	"back/appconfig/utils/define"
	"strings"
	"back/appconfig/utils"
)

// @APIVersion 1.0.0
// @Description beego Test API
// 使用注释路由
// @SecurityDefinition jwt apiKey Authorization header header
func init() {
	// 基础初始化
	BaseInit()

	beego.Router("/", &default_service.DefaultController{}, "*:ApiGetAll")
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

func BaseInit() {
	sentry.Init() // sentry
	models.Init() // 模型

	corsFilter()
	authFilter()
}
func corsFilter() {
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}

func authFilter() {
	var FilterUser = func(ctx *context.Context) {
		et := utils.EasyToken{}
		authtoken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))

		valid, _ := et.ValidateToken(authtoken, 0)
		if !valid {
			ctx.Redirect(302, "/view/user/login.html")
		}

		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}

	beego.InsertFilter("/api/*", beego.BeforeRouter, FilterUser)
}
