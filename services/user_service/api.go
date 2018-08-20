package user_service

import (
	"fmt"
	"back/appconfig/models"
	"back/appconfig/utils"

	"strings"
	. "back/appconfig/services/base_service"
)

var (
	ErrPhoneIsRegis     = ErrResponse{422001, "手机用户已经注册"}
	ErrUsernameIsRegis  = ErrResponse{422002, "用户名已经被注册"}
	ErrUsernameOrPasswd = ErrResponse{422003, "账号或密码错误。"}
)

type UserController struct {
	BaseController
}
type LoginToken struct {
	User  models.StaffUser `json:"user"`
	Token string           `json:"token"`
}

// @Title 注册新用户
// @Description 用户注册
// @Param	username	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {object} models.StaffUser
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (this *UserController) ApiRegister() {
	username := this.GetString("username")
	password := this.GetString("password")
	//code := this.GetString("code")

	if this.validateRegister(username, password, "") != nil {
		return
	}

	if models.CheckUserUsername(username) {
		this.WriteJsonWithCode(403, ErrUsernameIsRegis)
		return
	}

	password = utils.TransPassword(password) // 存储密码hash值
	user := models.StaffUser{
		Username: username,
		Password: password,
	}
	this.WriteJson(Response{0, "success.", models.CreateUser(user)})

}

// @Title 登录
// @Description 账号登录接口
// @Param	username	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /login [post]
func (this *UserController) ApiLogin() {
	//defaultController := default_service.DefaultController{this.BaseController}  // other controller
	//defaultController.GetAllPublic()

	this.loginTest() // 测试登陆日志
	this.LoginTest() // 测试登陆日志

	username := this.GetString("username")
	password := this.GetString("password")
	this.GetLogger().Msg("this is a message with trace id")
	this.GetLogger().Msgf("username: %s try to login.", username)

	user, ok := models.CheckUserAuth(username, password)
	if !ok {
		this.WriteJsonWithCode(403, ErrUsernameOrPasswd)
		return
	}

	et := utils.EasyToken{
		Username: user.Username,
		Uid:      user.Id,
		Expires:  utils.GetExpireTime(),
	}

	token, err := et.GetToken()
	if token == "" || err != nil {
		this.WriteJson(ErrResponse{0, err})
	} else {
		this.SetSession("uid", user.Id)
		this.WriteJson(Response{0, "success.", LoginToken{user, token}})
	}

}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {string}
// @Failure 401 unauthorized
// @router /auth [get]
// @Security jwt
func (this *UserController) ApiAuth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	valid, err := et.ValidateToken(authtoken, 0)
	if !valid {
		this.WriteJsonWithCode(401, ErrResponse{-1, fmt.Sprintf("%s", err)})
		return
	}
	this.WriteJson(Response{0, "success.", "is login"})

}
