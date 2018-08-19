package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"back/appconfig/utils"
	"github.com/astaxie/beego/logs"
)

type StaffUser struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Email    string    `json:"email" orm:"column(email);unique;size(50)"`
	Username string    `json:"username" orm:"column(username);unique;size(100);"`
	Password string    `json:"-" orm:"column(password);size(300)"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

// 多字段索引
func (u *StaffUser) TableIndex() [][]string {
	return [][]string{
		{"Id", "Username"},
	}
}

// 多字段唯一键
func (u *StaffUser) TableUnique() [][]string {
	return [][]string{
		{"Username", "Email"},
	}
}

func (u *StaffUser) TableName() string {
	return TableName("user")
}
func init() {
	orm.RegisterModel(new(StaffUser))
}

func Users() orm.QuerySeter {
	return OrmManager().QueryTable(new(StaffUser))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := Users().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserUsername(username string) bool {
	exist := Users().Filter("username", username).Exist()
	return exist
}

//创建用户
func CreateUser(user StaffUser) int64 {
	id, err := OrmManager().Insert(&user)
	if err != nil {
		logs.Error("数据库错误: 创建用户失败, user: %v, error: %v", user, err)
		return -1
	}
	return id
}

//检测手机和昵称是否注册
func CheckUserPhoneOrUsername(phone string, username string) bool {
	cond := orm.NewCondition()
	count, _ := Users().SetCond(cond.And("phone", phone).Or("username", username)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(username string, password string) (StaffUser, bool) {

	var user StaffUser

	err := Users().Filter("username", username).One(&user)

	if err != nil || user.Password != utils.TransPassword(password) {
		return user, false
	}
	return user, true
}
