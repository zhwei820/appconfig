package base_service

import (
	"github.com/zhwei820/appconfig/utils"
	"github.com/zhwei820/appconfig/models"
	"strings"
)
// 当前用户
func (this *BaseController) GetUser() (models.StaffUser, error) {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))

	uid := et.GetUserId(authtoken)

	var staffuser models.StaffUser
	err := models.StaffUsers().Filter("id", uid).One(&staffuser)

	return staffuser, err

}