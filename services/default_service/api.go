package default_service

import (
	. "back/appconfig/services/base_service"
	"github.com/astaxie/beego/validation"
	"strings"
	"back/appconfig/utils"
)

// Operations about object
type DefaultController struct {
	BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {string}
// @router / [get]
func (this *DefaultController) ApiGetAll() {
	this.WriteJson(Response{0, "success.", "API 1.0"})
}

type TestController struct {
	BaseController
}

// Address houses a users address information
type Address struct {
	Street string `valid:"Required"`
	City   string `valid:"Required"`
	Planet string `valid:"Required"`
	Phone  string `valid:"Required;MinSize(10)"`
}

// 验证函数写在 "valid" tag 的标签里d
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名
type User struct {
	Id        int
	Name      string     `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Age       int        `valid:"Range(1, 140)"`            // 1 <= Age <= 140，超出此范围即为不合法
	Email     string     `valid:"Email; MaxSize(100)"`      // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile    string     `valid:"Mobile"`                   // Mobile 必须为正确的手机号
	IP        string     `valid:"IP"`                       // IP 必须为一个正确的 IPv4 地址
	Addresses []*Address `valid:"Required"`                 // a person can have a home and cottage...
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

// @Title PostTest
// @Description PostTest
// @Param   body  body default_service.User 用户信息
// @Success 200 {string}
// @router / [post]
func (this *TestController) ApiPostTest() {

	this.GetLogger().Msg("api post test")
	this.GetLogger().Msg("api post test")
	this.GetLogger().Msg("api post test")
	this.GetLogger().Msg("api post test")
	this.GetLogger().Msg("api post test")
	this.GetLogger().Msg("api post test")

	var user User
	err := this.GetJson(&user)
	if err == nil {
		for _, addr := range user.Addresses {
			err = utils.Validate(addr) // 根据struct tag验证
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		this.WriteJsonWithCode(403, err.Error())
		return
	}
	this.WriteJson(user)
}
