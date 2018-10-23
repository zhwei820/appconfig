package auth_service

import (
	"github.com/zhwei820/appconfig/models"
	"github.com/zhwei820/appconfig/utils"
	"strings"
	. "github.com/zhwei820/appconfig/services/base_service"
	. "github.com/zhwei820/appconfig/utils/error_define"
	"github.com/zhwei820/appconfig/services/staffuser_service"
	"github.com/zhwei820/appconfig/utils/redisp"
	"github.com/bitly/go-simplejson"
	"github.com/rs/zerolog/log"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
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
	token, err := genTokenAndSave(user) // 获取token

	if err != nil {
		this.WriteJsonWithStatusCode(500, ErrorUnkown.Code, err.Error())
		return
	}

	this.WriteJson(LoginToken{user, token})
}

// @Summary 登出
// @Description 账号登出接口
// @Success 200 {string}
// @Failure 401 No Admin
// @router /api_logout [post]
// @Security jwt
func (this *AuthController) ApiLogout() {
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	c := redisp.CachePool.Get() // set redis session
	_, err := c.Do("DEL", authtoken)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("log out error:  %s", err.Error()))
	}
	this.WriteJson("logout")

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

	res := this.getTokenData(authtoken)
	if len(res) == 0 {
		return // 未授权
	}

	valid, err := et.ValidateToken(authtoken, 0)
	if !valid {
		this.WriteJsonWithStatusCode(401, ErrorUnkown.Code, err.Error())
		return
	}
	this.WriteJson("is login")

}

// @Summary 刷新token
// @Description 刷新token
// @Success 200 {string}
// @Failure 401 unauthorized
// @router /token_refresh [get]
// @Security jwt
func (this *AuthController) ApiTokenRefresh() {
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	res := this.getTokenData(authtoken)
	if len(res) == 0 {
		return // 未授权
	}

	c := redisp.CachePool.Get() // set redis session
	_, err := c.Do("DEL", authtoken)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("log out error:  %s", err.Error()))
	}

	var user models.StaffUser
	json.Unmarshal(res, &user)

	token, err := genTokenAndSave(user) // 获取token
	if err != nil {
		log.Error().Msg(fmt.Sprintf("log out error:  %s", err.Error()))
	}
	this.WriteJson(LoginToken{user, token})

}

// 根据token获取用户信息
func (this *AuthController) getTokenData(authtoken string) []byte {
	c := redisp.CachePool.Get() // get redis session
	res, err := redis.Bytes(c.Do("GET", authtoken))
	if err != nil {
		log.Error().Msg(err.Error())
	}
	if len(res) == 0 {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Ctx.WriteString("未授权")
	}
	return res

}
