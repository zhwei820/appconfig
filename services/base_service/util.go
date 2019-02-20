package base_service

import (
	"github.com/zhwei820/natsmicro/utils"
	"strings"
)

// 当前用户
func (this *BaseController) GetUser() (int64) {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	return et.GetUserId(authtoken)
}
