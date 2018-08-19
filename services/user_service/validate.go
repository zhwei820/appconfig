package user_service

import (
	"github.com/astaxie/beego/validation"
	. "back/appconfig/services/base_service"
)

func (this *UserController) validateRegister(phone, username, password, email string) error {

	valid := validation.Validation{}
	//表单验证
	valid.Required(phone, "phone").Message("手机必填")
	valid.Required(username, "username").Message("用户昵称必填")
	valid.Required(password, "password").Message("密码必填")
	valid.Mobile(phone, "phone").Message("手机号码不正确")
	valid.MinSize(username, 2, "username").Message("用户名最小长度为 2")
	valid.MaxSize(username, 40, "username").Message("用户名最大长度为 40")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.WriteJsonWithCode(403, ErrResponse{403001, map[string]string{err.Key: err.Message}})
			return err
		}
	}
	return nil

}
