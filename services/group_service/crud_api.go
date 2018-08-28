package group_service

import (
	"back/appconfig/models"
	. "back/appconfig/services/base_service"
	"strconv"
	"github.com/astaxie/beego"
	. "back/appconfig/utils/error_define"
)

type GroupController struct {
	BaseController
}

// @Summary 用户组列表
// @Description 用户组列
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @router /group [get]
func (this *GroupController) ApiGroupList() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var groups [] models.Group
	models.Groups().Limit(page_size, (page-1)*page_size).All(&groups)
	count, _ := models.Groups().Count()
	beego.Trace(count)

	this.WriteListJson(count, groups)
}

// @Summary 用户组
// @Description 用户组
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /group/:id [get]
func (this *GroupController) ApiDetailGroup() {
	id := this.Ctx.Input.Param(":id")

	var group models.Group
	err := models.Groups().Filter("id", id).One(&group)

	if err != nil {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(group)
}

// @Summary 新建用户组
// @Description 新建用户组
// @Param   body  body models.Group 用户组
// @Success 200 {string}
// @router /group [post]
func (this *GroupController) ApiCreateGroup() {

	var group models.Group
	err := this.GetJson(&group)
	if err != nil {
		this.WriteJsonWithStatusCode(403, err.Error())
		return
	}
	id, err := models.OrmManager().Insert(&group)
	if err != nil {
		this.WriteJsonWithStatusCode(403, err.Error())
		return
	}

	this.WriteJson(Response{0, "success.", id})
}

// @Summary 更新用户组
// @Description 更新用户组 (支持部分更新)
// @Param	id		path 	integer	true		"id"
// @Param   body  body models.Group 用户组
// @Success 200 {string}
// @router /group/:id [put]
func (this *GroupController) ApiUpdateGroup() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.Groups().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, "obj not exist!")
		return
	}

	var group models.Group
	keys, err := this.GetJsonWithKeys(&group)
	if err != nil {
		this.WriteJsonWithStatusCode(403, err.Error())
		return
	}
	group.Id = int64(id)
	num, err := models.OrmManager().Update(&group, keys...)
	if err != nil {
		this.WriteJsonWithStatusCode(403, err.Error())
		return
	}

	this.WriteJson(Response{0, "success.", num})
}

// @Summary 删除用户组
// @Description 删除用户组
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /group/:id [delete]
func (this *GroupController) ApiDeleteGroup() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.Groups().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, "obj not exist!")
		return
	}

	group := models.Group{Id: int64(id)}
	num, err := models.OrmManager().Delete(&group)

	if err != nil {
		this.WriteJsonWithStatusCode(403, err.Error())
		return
	}

	this.WriteJson(Response{0, "success.", num})
}
