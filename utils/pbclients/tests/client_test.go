package tests

import (
	"testing"
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	"github.com/davecgh/go-spew/spew"
	"github.com/zhwei820/appconfig/utils/rpc_register"
	"time"
	. "github.com/zhwei820/appconfig/pb/appconfig"
	. "github.com/zhwei820/appconfig/utils/pbclients/demo"
	"github.com/astaxie/beego"
)

func TestDiscoveryRegister(t *testing.T) {
	go func() {
		beego.Run()
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			data := SayInput{Query: "sdfdsf"}
			input, _ := data.Marshal()

			res, _ := CallRpc(DemoComsumer, "Hello", input)
			var out SayOutput
			out.Unmarshal(res)
			spew.Dump(out)
		}
	}
	rpc_register.Cancel()
}
