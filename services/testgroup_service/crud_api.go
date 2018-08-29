package testgroup_service

import (
	"back/appconfig/models"
	. "back/appconfig/services/base_service"
	"strconv"
	"github.com/astaxie/beego"
	. "back/appconfig/utils/error_define"
)

type TestGroupController struct {
	BaseController
}

// @Summary TestGroup列表
// @Description TestGroup列表
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @router /testgroup [get]
func (this *TestGroupController) ApiTestGroupList() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var testgroups [] models.TestGroup
	models.TestGroups().Limit(page_size, (page-1)*page_size).All(&testgroups)
	count, _ := models.TestGroups().Count()
	beego.Trace(count)

	this.WriteListJson(count, testgroups)
}

// @Summary TestGroup
// @Description TestGroup
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /testgroup/:id [get]
func (this *TestGroupController) ApiDetailTestGroup() {
	id := this.Ctx.Input.Param(":id")

	var testgroup models.TestGroup
	err := models.TestGroups().Filter("id", id).One(&testgroup)

	if err != nil {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(testgroup)
}

// @Summary 新建TestGroup
// @Description 新建TestGroup
// @Param   body  body models.TestGroup TestGroup
// @Success 200 {string}
// @router /testgroup [post]
func (this *TestGroupController) ApiCreateTestGroup() {

	var testgroup models.TestGroup
	err := this.GetJson(&testgroup)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	id, err := models.OrmManager().Insert(&testgroup)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(id)

}

// @Summary 更新TestGroup
// @Description 更新TestGroup (支持部分更新)
// @Param	id		path 	integer	true		"id"
// @Param   body  body models.TestGroup TestGroup
// @Success 200 {string}
// @router /testgroup/:id [put]
func (this *TestGroupController) ApiUpdateTestGroup() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.TestGroups().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	var testgroup models.TestGroup
	keys, err := this.GetJsonWithKeys(&testgroup)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	testgroup.Id = int64(id)
	num, err := models.OrmManager().Update(&testgroup, keys...)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}

// @Summary 删除TestGroup
// @Description 删除TestGroup
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /testgroup/:id [delete]
func (this *TestGroupController) ApiDeleteTestGroup() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.TestGroups().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	testgroup := models.TestGroup{Id: int64(id)}
	num, err := models.OrmManager().Delete(&testgroup)

	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}