package tests

import (
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	. "github.com/zhwei820/appconfig/pb/appconfig/sing"
	. "github.com/zhwei820/appconfig/utils/pbclients"
	"log"
	"strings"
	"errors"
	"reflect"
)

// deprecate
func CallRpc1(demoComsumer *Consumer, methodName string, input []byte) ([]byte, error) {
	out := []byte{}
	for ii := 0; ii < 2; ii++ {
		svc, _ := demoComsumer.GetService() // 远程方法实例
		if svc == nil {
			println("svc nil")
			return out, errors.New("svc nil")
		}

		ret := CallInstanceFunc(*(svc.(*SingService)), methodName, input)
		out = ret[0].Bytes()
		err := ret[1].Interface().(error)

		if err != nil {
			if strings.Contains(err.Error(), "connect") {
				demoComsumer.RemoveService(demoComsumer.Idx)
				continue // 连不上rpc服务器, 切换服务器节点, 尝试2次
			}
			log.Printf("%v", err)
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
	method := reflect.ValueOf(c).FieldByName(name)
	if method.IsValid() {
		result = method.Call(in)
	}

	return result
}
