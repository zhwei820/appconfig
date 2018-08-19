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
		Str("uid", this.GetSessionStr("uid"))
}

// 获取字符串型session
func (this *BaseController) GetSessionStr(key string) string {
	uid := this.GetSession(key)
	if uid == nil {
		return ""
	}
	return uid.(string)
}

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

func (this *BaseController) WriteJson(jsonData interface{}) {
	this.Data["json"] = jsonData
	this.ServeJSON()
}

func (this *BaseController) WriteJsonWithCode(code int, jsonData interface{}) {
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Data["json"] = jsonData
	this.ServeJSON()
}

//Response 结构体
type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  interface{} `json:"errmsg"`
}
