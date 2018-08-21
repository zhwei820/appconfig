package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiAuth",
			Router: `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiRegister",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiUpdateUser",
			Router: `/user`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiDeleteUser",
			Router: `/user`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/user_service:UserController"],
		beego.ControllerComments{
			Method: "ApiUserList",
			Router: `/user_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
