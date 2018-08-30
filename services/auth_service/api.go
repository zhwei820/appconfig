package auth_service

import (
	"back/appconfig/models"
	"back/appconfig/utils"
	"strings"
	. "back/appconfig/services/base_service"
	. "back/appconfig/utils/error_define"
	"back/appconfig/services/staffuser_service"
)

var (
	ErrUsernameOrPasswd = ErrStruct{422003, "账号或密码错误。"}
)

type AuthController struct {
	BaseController
}

// @Summary 登录
// @Description 账号登录接口
// @Param	username	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /login [post]
func (this *AuthController) ApiLogin() {

	username := this.GetString("username")
	password := this.GetString("password")
	this.GetLogger().Msg("this is a message with trace id")
	this.GetLogger().Msgf("username: %s try to login.", username)

	user, ok := models.CheckUserAuth(username, password)
	if !ok {
		this.WriteJsonWithStatusCode(403, ErrUsernameOrPasswd.Code, ErrUsernameOrPasswd.Msg)
		return
	}

	this.SetSession("uid", user.Id)
	this.Redirect("/view/admin/index2.html", 302)

}

// @Summary 注册
// @Description 账号注册接口
// @Param   body  body models.StaffUser StaffUser
// @Success 200 {string}
// @Failure 401 No Admin
// @router /register [post]
func (this *AuthController) ApiRegister() {
	defaultController := staffuser_service.StaffUserController{this.BaseController} // other controller
	defaultController.ApiCreateStaffUser()

}

// @Summary 认证测试
// @Description 测试错误码
// @Success 200 {string}
// @Failure 401 unauthorized
// @router /auth [get]
// @Security jwt
func (this *AuthController) ApiAuth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	valid, err := et.ValidateToken(authtoken, 0)
	if !valid {
		this.WriteJsonWithStatusCode(401, ErrorUnkown.Code, err.Error())
		return
	}
	this.WriteJson("is login")

}
