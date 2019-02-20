package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"time"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"path/filepath"
	"strings"
)

func init() {
	fname := beego.AppConfig.String("log_name")
	initLogRotate(fname)
	initLogger(fname)
}

func makeLogDir(fname string) {
	sl := strings.Split(fname, "/")
	path := filepath.Join(sl[0 : len(sl)-1]...)
	err := os.MkdirAll(path, 0777)
	if err != nil {
		logs.Error("mkdir error %v; %v", path, err)
	}

}

var opFile io.WriteCloser
var ticker *time.Ticker

func initLogger(fname string) {
	makeLogDir(fname)
	zerolog.TimestampFunc = func() time.Time { return time.Now().Round(time.Second) }
	refreshLogger(fname)
}

func refreshLogger(fname string) {

	var err error
	opFile, err = os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0660))
	if err != nil {
		fmt.Printf("%v", err)
	}

	log.Logger = zerolog.New(opFile).With().Timestamp().Logger()
}

func initLogRotate(fname string) {
	ticker = time.NewTicker(time.Second * 33)

	go func() {
		for t := range ticker.C {
			fmt.Printf("log rotate at %s\n", t)
			NewLogrotate(fname, 0, 10, 100, false, refreshLogger) //
		}
	}()

}

// 清除未释放资源
func Destroy() {
	if err := opFile.Close(); err != nil {
		opFile.Close()
	}
	ticker.Stop()

}
