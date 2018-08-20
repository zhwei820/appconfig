package main

import (
	_ "github.com/astaxie/beego/session/redis"

	"back/appconfig/models"
	"back/appconfig/utils/define"

	_ "back/appconfig/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	"time"
	"github.com/satori/go.uuid"
	"back/appconfig/utils/sentry"
	"github.com/rs/zerolog/log"
	_ "back/appconfig/utils/util"
	"back/appconfig/utils/util"
)

func init() {
	sentry.Init() // sentry
	models.Init() // 模型
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
		ctx.Request.Header.Add(define.TraceId, uuid.NewV4().String()) // trace id
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
}

func main() {

	defer func() {
		destroy()
	}()
	logTest()
	debug, _ := beego.AppConfig.Bool("debug")

	beego.BConfig.WebConfig.StaticDir["/asset"] = "./views/templates"

	if debug {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
		orm.Debug = true
	}
	orm.DefaultTimeLoc = time.UTC
	//beego.ErrorController(&base_service.ErrorController{})
	beego.BConfig.ServerName = "snail server 1.0"

	beego.Run()
}

func logTest() {
	t1 := time.Now()
	for ii := 0; ii < 100000; ii++ {
		log.Info().Msg("xxxxx")
	}
	t2 := time.Since(t1).Nanoseconds()

	println(float64(t2) / float64(time.Second))

}

func destroy() {
	if err := util.OpFile.Close(); err != nil {
		util.OpFile.Close()

	}
}
