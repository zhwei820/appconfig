package routers

import (
	"github.com/zhwei820/appconfig/utils/sentry"
	"github.com/zhwei820/appconfig/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/zhwei820/appconfig/utils/define"
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
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
}

func authFilter() {
	var FilterUser = func(ctx *context.Context) {
		test := beego.AppConfig.String("test")
		if test != "true" { // test 不验证token

			// // token base validation
			//	et := utils.EasyToken{}
			//	authtoken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))
			//
			//	valid, _ := et.ValidateToken(authtoken, 0)
			//	if !valid {
			//		ctx.Redirect(302, "/view/auth/login.html")
			//	}

			// session base validation
			uid := ctx.Input.CruSession.Get("uid")
			if uid == nil {
				ctx.Redirect(302, "/view/auth/login.html")
			}
		}
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}

	beego.InsertFilter("/api/*", beego.BeforeRouter, FilterUser)
}
