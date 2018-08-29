package user_service_c

import (
	"back/appconfig/models"
	"back/appconfig/utils"

	"strings"
	. "back/appconfig/services/base_service"
	. "back/appconfig/utils/error_define"
)

var (
	ErrPhoneIsRegis     = ErrStruct{422001, "手机用户已经注册"}
	ErrUsernameIsRegis  = ErrStruct{422002, "用户名已经被注册"}
	ErrUsernameOrPasswd = ErrStruct{422003, "账号或密码错误。"}
)

type UserController struct {
	BaseController
}
type LoginToken struct {
	User  models.StaffUser `json:"user"`
	Token string           `json:"token"`
}

// @Summary 注册新用户
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

	if this.validateRegister(username, password, "") != nil {
		return
	}

	if models.CheckUserUsername(username) {
		this.WriteJsonWithStatusCode(403, ErrUsernameIsRegis.Code, ErrUsernameIsRegis.Msg)
		return
	}

	password = utils.TransPassword(password) // 存储密码hash值
	user := models.StaffUser{
		Username: username,
		Password: password,
	}
	this.WriteJson(models.CreateUser(user))

}

// @Summary 更新用户
// @Description 更新用户
// @Param   body  body models.StaffUser 用户信息
// @Success 200 {object} models.StaffUser
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /user [put]
func (this *UserController) ApiUpdateUser() {

	var user models.StaffUser
	err := this.GetJson(&user)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	id, err := models.OrmManager().InsertOrUpdate(&user)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(id)
}

// @Summary 删除用户
// @Description 删除用户
// @Param	username	formData 	string	true		"用户昵称"
// @Success 200 {object} models.StaffUser
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /user [delete]
func (this *UserController) ApiDeleteUser() {
	var user models.StaffUser
	username := this.GetString("username")

	user = models.StaffUser{Username: username}
	id, err := models.OrmManager().Delete(&user)

	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(id)
}

// @Summary 登录
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
		this.WriteJsonWithStatusCode(403, ErrUsernameOrPasswd.Code, ErrUsernameOrPasswd.Msg)
		return
	}

	this.SetSession("uid", user.Id)
	this.Redirect("/view/admin/index2.html", 302)

}

// @Summary 认证测试
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
		this.WriteJsonWithStatusCode(401, ErrorUnkown.Code, err.Error())
		return
	}
	this.WriteJson("is login")

}

// @Summary 用户列表
// @Description 用户列表
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /user_list [get]
func (this *UserController) ApiUserList() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var users [] models.StaffUser
	models.StaffUsers().Limit(page_size, (page-1)*page_size).All(&users)
	count, _ := models.StaffUsers().Count()

	this.WriteListJson(count, users)
}
