package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"] = append(beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"],
		beego.ControllerComments{
			Method: "ApiTestGroupList",
			Router: `/testgroup`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"] = append(beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"],
		beego.ControllerComments{
			Method: "ApiCreateTestGroup",
			Router: `/testgroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"] = append(beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"],
		beego.ControllerComments{
			Method: "ApiDetailTestGroup",
			Router: `/testgroup/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"] = append(beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"],
		beego.ControllerComments{
			Method: "ApiUpdateTestGroup",
			Router: `/testgroup/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"] = append(beego.GlobalControllerRouter["back/appconfig/services/testgroup_service:TestGroupController"],
		beego.ControllerComments{
			Method: "ApiDeleteTestGroup",
			Router: `/testgroup/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
