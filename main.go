package main

import (
	_ "github.com/zhwei820/appconfig/routers"
	_ "github.com/zhwei820/appconfig/pbservice"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"time"
	"github.com/zhwei820/appconfig/utils/util"
)

func main() {

	defer func() {
		destroy() // 退出后清理资源
	}()
	debug, _ := beego.AppConfig.Bool("debug")

	beego.BConfig.WebConfig.StaticDir["/asset"] = "./views/templates"

	if debug {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
	}
	orm.DefaultTimeLoc = time.UTC
	beego.BConfig.ServerName = "snail server 1.0"

	beego.Run()

}

func destroy() {
	util.Destroy()

}
