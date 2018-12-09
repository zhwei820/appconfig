package utils

import (
	"github.com/astaxie/beego/validation"
	"errors"
)

// 根据struct tag验证字段
func Validate(ob interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(ob)
	if err != nil {
		return err
	}
	if !b {
		// validation does not pass
		msg := ""
		for _, err := range valid.Errors {
			msg += err.Key + ": " + err.Message + ";\n"
		}
		return errors.New(msg)
	}
	return nil
}
