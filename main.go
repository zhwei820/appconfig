package main

import (
	_ "github.com/zhwei820/appconfig/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/zhwei820/appconfig/utils/rpc_register"
	"github.com/zhwei820/appconfig/utils/util"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// https://stackoverflow.com/questions/41432193/how-to-delete-a-file-using-golang-on-program-exit/41455960#41455960
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // 程序退出清理资源
		<-sigs
		destroy()
		os.Exit(0)
	}()

	debug, _ := beego.AppConfig.Bool("debug")
	beego.BConfig.WebConfig.StaticDir["/asset"] = "./views/templates"
	if debug {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
	}
	orm.DefaultTimeLoc = time.UTC
	beego.BConfig.ServerName = "snail server 1.0"

	rpc_register.DiscoveryRegister() // DiscoveryRegister
	beego.Run()

}

func destroy() {
	log.Info().Msg("清除资源")
	util.Destroy()
	rpc_register.Cancel() // 取消注册 DiscoveryRegister
}
