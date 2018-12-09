package sentry

import (
	"github.com/getsentry/raven-go"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
	"net/http"
	"github.com/astaxie/beego/logs"
	"runtime"
	"github.com/astaxie/beego/context"
)

func Init() {
	raven.SetDSN(beego.AppConfig.String("sentry_dsn"))

	if !beego.BConfig.RecoverPanic {
		beego.BConfig.RecoverFunc = recoverPanic
	}
}

func recoverPanic(ctx *context.Context) {
	if err := recover(); err != nil {
		if err == beego.ErrAbort {
			return
		}
		if !beego.BConfig.RecoverPanic {
			errStr := fmt.Sprint(err)
			packet := raven.NewPacket(errStr, raven.NewException(errors.New(errStr),
				raven.NewStacktrace(2, 3, nil)), raven.NewHttp(ctx.Request))
			raven.Capture(packet, nil)
			ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		}

		var stack string
		logs.Critical("the request url is ", ctx.Input.URL())
		logs.Critical("Handler crashed with error", err)
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Critical(fmt.Sprintf("%s:%d", file, line))
			stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
		}

	}
}
