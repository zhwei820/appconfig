package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"],
		beego.ControllerComments{
			Method: "ApiStaffUserList",
			Router: `/staffuser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"],
		beego.ControllerComments{
			Method: "ApiCreateStaffUser",
			Router: `/staffuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"],
		beego.ControllerComments{
			Method: "ApiDetailStaffUser",
			Router: `/staffuser/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"],
		beego.ControllerComments{
			Method: "ApiUpdateStaffUser",
			Router: `/staffuser/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"] = append(beego.GlobalControllerRouter["back/appconfig/services/staffuser_service:StaffUserController"],
		beego.ControllerComments{
			Method: "ApiDeleteStaffUser",
			Router: `/staffuser/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
