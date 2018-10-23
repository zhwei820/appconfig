package routers

import (
	"github.com/zhwei820/appconfig/utils/sentry"
	"github.com/zhwei820/appconfig/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/zhwei820/appconfig/utils/define"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/context"
	"github.com/zhwei820/appconfig/utils"
	"strings"
	"github.com/zhwei820/appconfig/utils/redisp"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
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
		RequestCounter.Inc() // prom count stat

		test := beego.AppConfig.String("test")
		if test != "true" { // test 不验证token

			// // token base validation
			et := utils.EasyToken{}
			authtoken := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))

			valid, _ := et.ValidateToken(authtoken, 0)
			if !valid {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.WriteString("未授权")
				return
			}
			//userId := et.GetUserId(authtoken)
			//if userId == 0 {
			//	ctx.WriteString("未授权")
			//	return
			//}

			c := redisp.CachePool.Get() // get redis session
			res, err := redis.Bytes(c.Do("GET", authtoken))
			if err != nil {
				log.Error().Msg(err.Error())
			}
			if len(res) == 0 {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.WriteString("未授权")
				return
			}

			// // session base validation
			//uid := ctx.Input.CruSession.Get("uid")
			//if uid == nil {
			//	ctx.Redirect(302, "/view/auth/login.html")
			//}
		}
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}

	beego.InsertFilter("/api/*", beego.BeforeRouter, FilterUser)
}
