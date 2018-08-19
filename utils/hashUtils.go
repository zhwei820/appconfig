package utils

import (

	"crypto/sha256"

	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
)
//
//func Sha1(input string) string {
//	if input == "" {
//		return "adc83b19e793491b1c6ea0fd8b46cd9f32e592fc"
//	}
//	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
//}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TransPassword(password string) string {
	return Sha256(password + beego.AppConfig.String("jwt::secret"))
}

func Base64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
