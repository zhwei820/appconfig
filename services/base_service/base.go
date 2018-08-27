package base_service

import (
	"github.com/astaxie/beego"
	"back/appconfig/utils/define"
	"encoding/json"
	"back/appconfig/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) GetLogger() *zerolog.Event {
	// zero log with trace id
	return log.Info().Str("trace_id", this.Ctx.Request.Header.Get(define.TraceId)).
		Interface("uid", this.GetSession("uid"))
}

// 绑定json
func (this *BaseController) GetJson(ob interface{}) (error) {
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, ob); err == nil {
		err := utils.Validate(ob) // 根据struct tag验证
		if err != nil {
			return err
		}

		return nil
	}
	return err
}

// 绑定json, 并返回1st level key list
func (this *BaseController) GetJsonWithKeys(ob interface{}) ([] string, error) {
	var err = this.GetJson(ob)
	keys := utils.GetJsonKeys(this.Ctx.Input.RequestBody)
	return keys, err
}

// 输出json
func (this *BaseController) WriteJson(jsonData interface{}) {
	this.Data["json"] = jsonData
	this.ServeJSON()
}

// 输出带code的json
func (this *BaseController) WriteJsonWithCode(code int, jsonData interface{}) {
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Data["json"] = jsonData
	this.ServeJSON()
}

//Response 结构体
type Response struct {
	Errcode int         `json:"code"`
	Errmsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type ResponseList struct {
	Errcode int         `json:"code"`
	Errmsg  string      `json:"msg"`
	Count   int64       `json:"count"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"code"`
	Errmsg  interface{} `json:"msg"`
}
