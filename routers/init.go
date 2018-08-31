package routers

import (
	"back/appconfig/utils/sentry"
	"back/appconfig/models"
	"github.com/astaxie/beego"
	"back/appconfig/utils"
	"strings"
	"back/appconfig/utils/define"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/context"
)

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
		test := beego.AppConfig.String("test")
		if test != "true" { // test 不验证token
			et := utils.EasyToken{}
			authtoken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))

			valid, _ := et.ValidateToken(authtoken, 0)
			if !valid {
				ctx.Redirect(302, "/view/user/login.html")
			}
		}
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}

	beego.InsertFilter("/api/*", beego.BeforeRouter, FilterUser)
}
