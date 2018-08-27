package util

import (
	"time"
	"os"
	"io"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
	"path/filepath"
	"strings"
)

func init() {
	initLogger(beego.AppConfig.String("log_name"))
}

func makeLogDir(fname string) {
	sl := strings.Split(fname, "/")
	path := filepath.Join(sl[0 : len(sl)-1]...)
	err := os.MkdirAll(path, 0777)
	if err != nil {
		logs.Error("mkdir error; %v; %v", path, err)
	}

}

var OpFile io.WriteCloser

func initLogger(fname string) {
	makeLogDir(fname)

	opFile, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0660))
	if err != nil {
		fmt.Printf("%v", err)
	}

	OpFile = opFile

	zerolog.TimestampFunc = func() time.Time { return time.Now().Round(time.Second) }
	log.Logger = zerolog.New(OpFile).With().Timestamp().Logger()

	//wg.Add(10)
	//for ii := 0; ii < 10; ii ++ {
	//	go func() {
	//		for i := 0; i < count; i++ {
	//			log.Error().
	//				Int("Fault", 41650+i).Msg("Some Message")
	//		}
	//		wg.Done()
	//
	//	}()
	//}
	//wg.Wait()
}
