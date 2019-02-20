package base_service

import (
	"github.com/astaxie/beego"
	"github.com/zhwei820/natsmicro/utils/define"
	"encoding/json"
	"github.com/zhwei820/natsmicro/utils"
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

// 输出json list
func (this *BaseController) WriteListJson(count int64, Data interface{}) {
	jsonData := ResponseList{0, "success", count, Data}
	this.Data["json"] = jsonData
	this.ServeJSON()
}

// 输出json
func (this *BaseController) WriteJson(Data interface{}) {
	jsonData := Response{0, "success", Data}
	this.Data["json"] = jsonData
	this.ServeJSON()
}

// 输出带code的json
func (this *BaseController) WriteJsonWithStatusCode(statucCode int, errorCode int, errMsg string) {
	this.Ctx.ResponseWriter.WriteHeader(statucCode)
	jsonData := Response{errorCode, errMsg, nil}
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
