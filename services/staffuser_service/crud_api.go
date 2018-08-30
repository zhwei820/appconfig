package staffuser_service

import (
	"back/appconfig/models"
	. "back/appconfig/services/base_service"
	. "back/appconfig/utils/error_define"

	"strconv"
	"github.com/astaxie/beego"
	"back/appconfig/utils"

)

var (
	ErrPhoneIsRegis     = ErrStruct{422001, "手机用户已经注册"}
	ErrUsernameIsRegis  = ErrStruct{422002, "用户名已经被注册"}
)

type StaffUserController struct {
	BaseController
}

// @Summary StaffUser列表
// @Description StaffUser列表
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @router /staffuser [get]
func (this *StaffUserController) ApiStaffUserList() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var staffusers [] models.StaffUser
	models.StaffUsers().Limit(page_size, (page-1)*page_size).All(&staffusers)
	count, _ := models.StaffUsers().Count()
	beego.Trace(count)

	this.WriteListJson(count, staffusers)
}

// @Summary StaffUser
// @Description StaffUser
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /staffuser/:id [get]
func (this *StaffUserController) ApiDetailStaffUser() {
	id := this.Ctx.Input.Param(":id")

	var staffuser models.StaffUser
	err := models.StaffUsers().Filter("id", id).One(&staffuser)

	if err != nil {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(staffuser)
}

// @Summary 新建StaffUser
// @Description 新建StaffUser
// @Param   body  body models.StaffUser StaffUser
// @Success 200 {string}
// @router /staffuser [post]
func (this *StaffUserController) ApiCreateStaffUser() {

	var staffuser models.StaffUser
	err := this.GetJson(&staffuser)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}

	if this.validateRegister(staffuser.Username, staffuser.Password, "") != nil {
		return
	}

	if models.CheckUserUsername(staffuser.Username) {
		this.WriteJsonWithStatusCode(403, ErrUsernameIsRegis.Code, ErrUsernameIsRegis.Msg)
		return
	}

	staffuser.Password = utils.TransPassword(staffuser.Password) // 存储密码hash值

	id, err := models.OrmManager().Insert(&staffuser)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(id)

}

// @Summary 更新StaffUser
// @Description 更新StaffUser (支持部分更新)
// @Param	id		path 	integer	true		"id"
// @Param   body  body models.StaffUser StaffUser
// @Success 200 {string}
// @router /staffuser/:id [put]
func (this *StaffUserController) ApiUpdateStaffUser() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.StaffUsers().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	var staffuser models.StaffUser
	keys, err := this.GetJsonWithKeys(&staffuser)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	staffuser.Id = int64(id)
	num, err := models.OrmManager().Update(&staffuser, keys...)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}

// @Summary 删除StaffUser
// @Description 删除StaffUser
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /staffuser/:id [delete]
func (this *StaffUserController) ApiDeleteStaffUser() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.StaffUsers().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	staffuser := models.StaffUser{Id: int64(id)}
	num, err := models.OrmManager().Delete(&staffuser)

	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}