package say

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/zhwei820/appconfig/pbservice/pb/appconfig/say"
	"github.com/zhwei820/appconfig/pbservice/pbclients"
	"strings"
)

var SayClientRpc *pbclients.ClientRpcs

func CreateNewRpcClients(remoteSvcUrlsKey string) * pbclients.ClientRpcs {
	rpcUrls := beego.AppConfig.DefaultString(remoteSvcUrlsKey, "")
	rpcUrlList := strings.Split(rpcUrls, ";")
	clientRpc:= pbclients.CreateNewInstance()
	for _, rpcUrl := range rpcUrlList{
		rpclient := rpc.NewHTTPClient(rpcUrl)
		clientRpc.Clients = append(clientRpc.Clients, rpclient)

		var sayService *say.SayService // replace this
		rpclient.UseService(&sayService)
		clientRpc.Services = append(clientRpc.Services, sayService)
	}
	return clientRpc
}

func init()  {
	SayClientRpc = CreateNewRpcClients("rpc_urls")
}