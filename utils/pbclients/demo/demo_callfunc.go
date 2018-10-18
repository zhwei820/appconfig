package demo

import (
	"strings"
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	. "github.com/zhwei820/appconfig/pb/appconfig"
	"log"
	"errors"
)

func CallRpc(demoComsumer *DemoComsumerStruct, name string, input []byte) (res []byte, err error) {
	flag := 1
	for ii := 0; ii <= 1; ii++ {

		svc, _ := demoComsumer.GetService() // 远程方法实例
		if svc == nil {
			println("svc nil")
			return res, errors.New("svc nil")
		}
		res, err = funcMap[name](svc, input)

		if err != nil {
			if strings.Contains(err.Error(), "connect") {
				demoComsumer.RemoveService(demoComsumer.Idx)
				flag ++ // 连不上rpc服务器, 切换服务器节点, 尝试2次
				if flag <= 2 {
					continue
				}
			}
			log.Printf("%v", err)
		}
		return res, err
	}
	return res, err
}

func testCallHello(svc interface{}, input []byte) ([]byte, error) {
	res, err := svc.(*SingService).Hello(input)
	return res, err

}

var funcMap map[string]func(svc interface{}, input []byte) ([]byte, error)

func init() {
	funcMap = map[string]func(svc interface{}, input []byte) ([]byte, error){
		"Hello": testCallHello,
	}
}
