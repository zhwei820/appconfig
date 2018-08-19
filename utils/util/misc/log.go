package misc

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	path := filepath.Join("data", "log")
	err := os.MkdirAll(path, 0777)
	if err != nil {
		logs.Error("mkdir error; %v", err)
	}

	logName := beego.AppConfig.DefaultString("log_name", "example.log")
	logs.EnableFuncCallDepth(true)
	logs.Async(1e4)

	loglevel := 6 // info

	if beego.AppConfig.String("runmode") == "dev" {
		logs.SetLogger("console")
		loglevel = 7 // debug
	} else {
		beego.BeeLogger.DelLogger("console")
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"./data/log/%v","level":%v,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`, logName, loglevel))

}
