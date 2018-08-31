package default_service

import (
	. "github.com/zhwei820/appconfig/services/base_service"
)

// Operations about object
type DefaultController struct {
	BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {string}
// @router / [get]
func (this *DefaultController) ApiGetAll() {
	this.WriteJson(Response{0, "success.", "API 1.0"})
}

func (this *DefaultController) Html() {
	dir := this.Ctx.Input.Param(":dir")
	html := this.Ctx.Input.Param(":html")
	if html != "login" && html != "register" {
		this.Layout = "templates/base/userinfo.html"
	}

	this.TplName = "templates/" + dir + "/" + html + ".html"
}
