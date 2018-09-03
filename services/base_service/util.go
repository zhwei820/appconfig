package base_service

import (
	"github.com/zhwei820/appconfig/utils"
	"strings"
)
// 当前用户
func (this *BaseController) GetUser() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	valid, err := et.ValidateToken(authtoken, 0)
	if !valid {
		this.WriteJsonWithStatusCode(401, ErrorUnkown.Code, err.Error())
		return
	}
}