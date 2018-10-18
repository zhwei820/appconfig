package tests

import (
	"testing"
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	"github.com/zhwei820/appconfig/utils/rpc_register"
	"github.com/zhwei820/appconfig/pb/appconfig/sing"
	singclient "github.com/zhwei820/appconfig/utils/pbclients/sing"
	sayclient "github.com/zhwei820/appconfig/utils/pbclients/say"

	"github.com/astaxie/beego"
	"time"
	"github.com/davecgh/go-spew/spew"
	"github.com/zhwei820/appconfig/pb/appconfig/say"
)

func TestDiscoveryRegister(t *testing.T) {
	go func() {
		beego.Run()
	}()

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				data := sing.SingInput{Query: "sdfdsf"}
				input, _ := data.Marshal()

				res, _ := singclient.SingRpc(singclient.SingComsumer, "SingHello", input)
				var out sing.SingOutput
				out.Unmarshal(res)
				spew.Dump(out)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				data := say.SayInput{Query: "sdfdsf"}
				input, _ := data.Marshal()

				res, _ := sayclient.SayRpc(sayclient.SayComsumer, "SayHello", input)
				var out say.SayOutput
				out.Unmarshal(res)
				spew.Dump(out)
			}
		}
	}()
	select {}
	rpc_register.Cancel()
}
