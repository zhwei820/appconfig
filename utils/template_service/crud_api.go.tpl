package {{ model_data.obj }}_service

import (
	"back/appconfig/models"
	. "back/appconfig/services/base_service"
	"strconv"
	"github.com/astaxie/beego"
	. "back/appconfig/utils/error_define"
)

type {{ model_data.objBig }}Controller struct {
	BaseController
}

// @Summary {{ model_data.objBig }}列表
// @Description {{ model_data.objBig }}列表
// @Param	page	query 	string	true		"page"
// @Param	page_size	query 	string	true		"page_size"
// @Success 200 {string}
// @router /{{ model_data.obj }} [get]
func (this *{{ model_data.objBig }}Controller) Api{{ model_data.objBig }}List() {

	page, _ := this.GetInt("page")
	page_size, _ := this.GetInt("page_size")
	this.GetLogger().Msg("this is a message with trace id")
	var {{ model_data.obj }}s [] models.{{ model_data.objBig }}
	models.{{ model_data.objBig }}s().Limit(page_size, (page-1)*page_size).All(&{{ model_data.obj }}s)
	count, _ := models.{{ model_data.objBig }}s().Count()
	beego.Trace(count)

	this.WriteListJson(count, {{ model_data.obj }}s)
}

// @Summary {{ model_data.objBig }}
// @Description {{ model_data.objBig }}
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /{{ model_data.obj }}/:id [get]
func (this *{{ model_data.objBig }}Controller) ApiDetail{{ model_data.objBig }}() {
	id := this.Ctx.Input.Param(":id")

	var {{ model_data.obj }} models.{{ model_data.objBig }}
	err := models.{{ model_data.objBig }}s().Filter("id", id).One(&{{ model_data.obj }})

	if err != nil {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson({{ model_data.obj }})
}

// @Summary 新建{{ model_data.objBig }}
// @Description 新建{{ model_data.objBig }}
// @Param   body  body models.{{ model_data.objBig }} {{ model_data.objBig }}
// @Success 200 {string}
// @router /{{ model_data.obj }} [post]
func (this *{{ model_data.objBig }}Controller) ApiCreate{{ model_data.objBig }}() {

	var {{ model_data.obj }} models.{{ model_data.objBig }}
	err := this.GetJson(&{{ model_data.obj }})
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	id, err := models.OrmManager().Insert(&{{ model_data.obj }})
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}

	this.WriteJson(id)

}

// @Summary 更新{{ model_data.objBig }}
// @Description 更新{{ model_data.objBig }} (支持部分更新)
// @Param	id		path 	integer	true		"id"
// @Param   body  body models.{{ model_data.objBig }} {{ model_data.objBig }}
// @Success 200 {string}
// @router /{{ model_data.obj }}/:id [put]
func (this *{{ model_data.objBig }}Controller) ApiUpdate{{ model_data.objBig }}() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.{{ model_data.objBig }}s().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	var {{ model_data.obj }} models.{{ model_data.objBig }}
	keys, err := this.GetJsonWithKeys(&{{ model_data.obj }})
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorInvalidJSON.Code, err.Error())
		return
	}
	{{ model_data.obj }}.Id = int64(id)
	num, err := models.OrmManager().Update(&{{ model_data.obj }}, keys...)
	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}

// @Summary 删除{{ model_data.objBig }}
// @Description 删除{{ model_data.objBig }}
// @Param	id		path 	integer	true		"id"
// @Success 200 {string}
// @router /{{ model_data.obj }}/:id [delete]
func (this *{{ model_data.objBig }}Controller) ApiDelete{{ model_data.objBig }}() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	exist := models.{{ model_data.objBig }}s().Filter("id", id).Exist()
	if !exist {
		this.WriteJsonWithStatusCode(404, ErrorDatabase.Code, "obj not exist!")
		return
	}

	{{ model_data.obj }} := models.{{ model_data.objBig }}{Id: int64(id)}
	num, err := models.OrmManager().Delete(&{{ model_data.obj }})

	if err != nil {
		this.WriteJsonWithStatusCode(403, ErrorDatabase.Code, err.Error())
		return
	}
	this.WriteJson(num)

}
