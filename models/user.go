package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/zhwei820/natsmicro/utils"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strings"
)

type StaffUser struct {
	Id       int64  `json:"id" orm:"column(id);pk;auto;unique"`
	Username string `json:"username" orm:"column(username);unique;size(100);"`
	Password string `json:"password" orm:"column(password);size(300)"`
	Group    string `json:"group" orm:"column(group);size(100)"`

	Created time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
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
		{"Username"},
	}
}

func (u *StaffUser) TableName() string {
	return TableName("user")
}
func init() {
	orm.RegisterModel(new(StaffUser))
}

func StaffUsers() orm.QuerySeter {
	return OrmManager().QueryTable(new(StaffUser))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := StaffUsers().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserUsername(username string) bool {
	exist := StaffUsers().Filter("username", username).Exist()
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
	count, _ := StaffUsers().SetCond(cond.And("phone", phone).Or("username", username)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(username string, password string) (StaffUser, bool) {

	var user StaffUser

	err := StaffUsers().Filter("username", username).One(&user)

	if err != nil || user.Password != utils.TransPassword(password) {
		return user, false
	}
	return user, true
}

// 获取用户和权限
func GetUserAndPermission(uid int64) orm.Params {
	sql := "SELECT a.id, a.username, a.password, a.group, a.create_at, a.update_at, " +
		"b.permission FROM user a INNER JOIN `group` b ON a.group = b.group WHERE a.id = ?;"

	var maps []orm.Params
	num, err := OrmManager().Raw(sql, uid).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["username"]) // slene
	}
	if len(maps) > 0 {
		userinfo := maps[0]
		userinfo["permission"] = strings.Split(userinfo["permission"].(string), ",")
		return userinfo
	}

	return orm.Params{}
}
