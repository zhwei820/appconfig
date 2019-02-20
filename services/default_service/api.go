package default_service

import (
	. "github.com/zhwei820/natsmicro/services/base_service"
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
	if dir != "auth" { // 前端页面需要登陆
		uid := this.GetSession("uid")
		if uid == nil {
			this.Ctx.Redirect(302, "/view/auth/login.html")
			return
		}
	}

	html := this.Ctx.Input.Param(":html")
	if html != "login" && html != "register" {
		this.Layout = "templates/base/userinfo.html"
	}

	this.TplName = "templates/" + dir + "/" + html + ".html"
}
