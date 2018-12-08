package say

import (
	"github.com/zhwei820/appconfig/pbservice/pbclients"
	"strings"
	_ "github.com/zhwei820/appconfig/utils/gotests"
	_ "github.com/zhwei820/appconfig/routers"

	"log"
	"errors"
	"github.com/zhwei820/appconfig/pbservice/pb/appconfig/say"
)

func SayRpc(clientRpcs *pbclients.ClientRpcs, name string, input []byte) (res []byte, err error) {
	flag:=1
	for ii := 0; ii <= 1; ii++ {
		svc, idx, pos, _ := clientRpcs.GetService() // 远程方法实例
		if svc == nil {
			println("svc nil")
			return res, errors.New("svc nil")
		}
		res, err = funcMap[name](svc, input)

		if err != nil {
			if strings.Contains(err.Error(), "connect") {
				clientRpcs.KnockdownService(idx, pos)
				flag +=1
				if flag <=2{
					continue
				}
			}
			log.Printf("%v", err)
		}
		return res, err
	}
	return res, err
}

func CallSayHello(svc interface{}, input []byte) ([]byte, error) {
	res, err := svc.(*say.SayService).SayHello(input)
	return res, err
}

var funcMap map[string]func(svc interface{}, input []byte) ([]byte, error)

func init() {
	funcMap = map[string]func(svc interface{}, input []byte) ([]byte, error){
		"SayHello": CallSayHello,
	}
}
