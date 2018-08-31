package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"] = append(beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"],
		beego.ControllerComments{
			Method: "ApiGroupList",
			Router: `/group`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"] = append(beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"],
		beego.ControllerComments{
			Method: "ApiCreateGroup",
			Router: `/group`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"] = append(beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"],
		beego.ControllerComments{
			Method: "ApiDetailGroup",
			Router: `/group/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"] = append(beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"],
		beego.ControllerComments{
			Method: "ApiUpdateGroup",
			Router: `/group/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"] = append(beego.GlobalControllerRouter["github.com/zhwei820/appconfig/services/group_service:GroupController"],
		beego.ControllerComments{
			Method: "ApiDeleteGroup",
			Router: `/group/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
