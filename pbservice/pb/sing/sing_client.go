// Code generated by protoc-gen-hprose , DO NOT EDIT.
// source: sing.proto

package sing


import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/zhwei820/BasePbClient"
	"strings"
	"errors"
	"github.com/rs/zerolog/log"
)


var SingClientRpc *pbclients.ClientRpcs

func CreateNewSingRpcClients(remoteSvcUrlsKey string) * pbclients.ClientRpcs {
	rpcUrls := beego.AppConfig.DefaultString(remoteSvcUrlsKey, "")
	rpcUrlList := strings.Split(rpcUrls, ";")
	clientRpc:= pbclients.CreateNewInstance()
	for _, rpcUrl := range rpcUrlList{
		rpclient := rpc.NewHTTPClient(rpcUrl)
		clientRpc.Clients = append(clientRpc.Clients, rpclient)

		var singService *SingService // replace this
		rpclient.UseService(&singService)
		clientRpc.Services = append(clientRpc.Services, singService)
	}
	return clientRpc
}

func init()  {
	SingClientRpc = CreateNewSingRpcClients("singrpc_urls")
}


func SingRpc(clientRpcs *pbclients.ClientRpcs, name string, input []byte) (res []byte, err error) {
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

var funcMap map[string]func(svc interface{}, input []byte) ([]byte, error)




func CallSingHello(svc interface{}, input []byte) ([]byte, error) {
	res, err := svc.(*SingService).SingHello(input)
	return res, err
}



func init() {
	funcMap = map[string]func(svc interface{}, input []byte) ([]byte, error){
		
    "SingHello": CallSingHello,
	}
}
