package tests

import (
	"testing"
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	"github.com/davecgh/go-spew/spew"
	"github.com/zhwei820/appconfig/utils/rpc_register"
	"time"
	. "github.com/zhwei820/appconfig/pb/appconfig"
	. "github.com/zhwei820/appconfig/utils/pbclients"
	"log"
	"github.com/astaxie/beego"
	"strings"
)

func TestDiscoveryRegister(t *testing.T) {
	go func() {
		beego.Run()
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			svc, _ := DemoComsumer.GetService()
			if svc == nil {
				println("svc nil")
			}
			data := SayInput{Query: "sdfdsf"}
			input, _ := data.Marshal()
			res, err := svc.(*SingService).Hello(input)
			if err != nil {
				if strings.Contains(err.Error(), "connect"){
					DemoComsumer.RemoveService(DemoComsumer.Idx)
				}
				log.Printf("%v", err.Error())
				continue
			}
			var out SayOutput
			out.Unmarshal(res)
			spew.Dump(out)
		}
	}
	rpc_register.Cancel()
}
