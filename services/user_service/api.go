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
		this.WriteJsonWithCode(401, ErrResponse{-1, fmt.Sprintf("%s", err)})
		return
	}
	this.WriteJson(Response{0, "success.", "is login"})

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
	models.Users().Limit(page_size, (page-1)*page_size).All(&users)
	count, _ := models.Users().Count()

	this.WriteJson(ResponseList{0, "success.", count, users})
}

// @Summary 用户组列表
// @Description 用户组列表用户组列表用户组列表
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /group_list [get]
func (this *UserController) ApiGroupList() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var groups [] models.Group
	models.Groups().Limit(page_size, (page-1)*page_size).All(&groups)
	count, _ := models.Groups().Count()

	this.WriteJson(ResponseList{0, "success.", count, groups})
}

// @Summary 新建用户组
// @Description 新建用户组
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /usergroup [post]
func (this *UserController) ApiCreateOrUpdateGroup() {

	var group models.Group
	err := this.GetJson(&group)
	if err != nil {
		this.WriteJsonWithCode(403, err.Error())
		return
	}
	id, err := models.OrmManager().InsertOrUpdate(&group)
	if err != nil {
		this.WriteJsonWithCode(403, err.Error())
		return
	}

	this.WriteJson(Response{0, "success.", id})
}

// @Summary 删除用户组
// @Description 删除用户组
// @Param	groupname	query 	string	true		"groupname"
// @Success 200 {string}
// @Failure 401 No Admin
// @router /usergroup [delete]
func (this *UserController) ApiDeleteGroup() {

	var group models.Group
	groupname := this.GetString("groupname")

	group = models.Group{GroupName: groupname}
	id, err := models.OrmManager().Delete(&group)

	if err != nil {
		this.WriteJsonWithCode(403, err.Error())
		return
	}

	this.WriteJson(Response{0, "success.", id})
}
