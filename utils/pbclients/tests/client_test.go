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
	"reflect"
	"errors"
)

func CallRpc(demoComsumer *Consumer, methodName string, input []byte) (out []byte, err error) {
	out = []byte{}
	for ii := 0; ii < 2; ii++ {
		svc, _ := demoComsumer.GetService() // 远程方法实例
		if svc == nil {
			println("svc nil")
			return out, errors.New("svc nil")
		}

		ret := CallInstanceFunc(*(svc.(*SingService)), methodName, input)
		out = ret[0].Bytes()
		er := ret[1].Interface()

		if err != nil {
			if strings.Contains(err.Error(), "connect") {
				demoComsumer.RemoveService(demoComsumer.Idx)
				continue // 连不上rpc服务器, 切换服务器节点, 尝试2次
			}
			log.Printf("%v", er)
			return out, err
		}
		return out, nil
	}
	return out, nil

}

func CallInstanceFunc(c interface{}, name string, params ... interface{}) (result []reflect.Value) {

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	method := reflect.ValueOf(c).FieldByName("Hello")
	if method.IsValid() {
		result = method.Call(in)
	}

	return result
}

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
