package auth_service

import (
	"github.com/zhwei820/appconfig/models"
	"github.com/zhwei820/appconfig/utils"
	"strings"
	. "github.com/zhwei820/appconfig/services/base_service"
	. "github.com/zhwei820/appconfig/utils/error_define"
	"github.com/zhwei820/appconfig/services/staffuser_service"
	"github.com/bitly/go-simplejson"
)

var (
	ErrUsernameOrPasswd = ErrStruct{422003, "账号或密码错误。"}
)

type AuthController struct {
	BaseController
}

type LoginToken struct {
	User  models.StaffUser `json:"user"`
	Token string           `json:"token"`
}

// @Summary 登录
// @Description 账号登录接口
// @Param	username	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /login [post]
func (this *AuthController) SessionLogin() {

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

// @Summary 登出
// @Description 账号登出接口
// @Success 200 {string}
// @Failure 401 No Admin
// @router /logout [post]
func (this *AuthController) SessionLogout() {

	this.DelSession("uid")
	this.Redirect("/view/auth/login.html", 302)
}

// @Summary 登录
// @Description 账号登录接口
// @Param   body  body extra.UserRegister UserRegister
// @Success 200 {string}
// @Failure 401 No Admin
// @router /api_login [post]
func (this *AuthController) ApiLogin() {
	j, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	username := j.Get("username").MustString()
	password := j.Get("password").MustString()

	this.GetLogger().Msg("this is a message with trace id")
	this.GetLogger().Msgf("username: %s try to login.", username)

	user, ok := models.CheckUserAuth(username, password)
	if !ok {
		this.WriteJsonWithStatusCode(403, ErrUsernameOrPasswd.Code, ErrUsernameOrPasswd.Msg)
		return
	}
	et := utils.EasyToken{
		Uid:     user.Id,
		Expires: utils.GetExpireTime(),
	}
	token, err := et.GetToken()
	if err != nil {
		this.WriteJsonWithStatusCode(500, ErrorUnkown.Code, err.Error())
		return
	}
	this.WriteJson(LoginToken{user, token})
}

// @Summary 注册
// @Description 账号注册接口
// @Param   body  body extra.UserRegister UserRegister
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
