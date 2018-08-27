package routers

import (
	"github.com/astaxie/beego"
	"back/appconfig/services/user_service"
	"back/appconfig/services/default_service"
	"back/appconfig/services/group_service"
	"github.com/satori/go.uuid"

	_ "github.com/astaxie/beego/session/redis"

	"back/appconfig/models"
	"back/appconfig/utils/define"

	"github.com/astaxie/beego/context"
	"back/appconfig/utils/sentry"
	_ "back/appconfig/utils/util"
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
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user_service.UserController{},
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
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}
