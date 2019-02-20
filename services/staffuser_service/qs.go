package staffuser_service

import (
	"github.com/astaxie/beego/orm"
	"github.com/zhwei820/natsmicro/models"
)
// 用户列表筛选
func (this *StaffUserController) ListFilterQs() orm.QuerySeter {
	qs := models.StaffUsers()
	username__icontains := this.GetString("username__icontains") // 用户名
	if username__icontains != "" {
		qs = qs.Filter("username__icontains", username__icontains)
	}
	return qs
}
