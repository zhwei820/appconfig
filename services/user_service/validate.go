package user_service

import (
	"github.com/astaxie/beego/validation"
	. "back/appconfig/utils/error_define"
)

func (this *UserController) validateRegister(username, password, email string) error {

	valid := validation.Validation{}
	//表单验证
	valid.Required(username, "username").Message("用户昵称必填")
	valid.Required(password, "password").Message("密码必填")
	valid.MinSize(username, 2, "username").Message("用户名最小长度为 2")
	valid.MaxSize(username, 40, "username").Message("用户名最大长度为 40")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.WriteJsonWithStatusCode(403, ErrValidation.Code, err.Key+": "+err.Message)
			return err
		}
	}
	return nil

}
