package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Group struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Email    string    `json:"email" orm:"column(email);unique;size(50)"`
	Username string    `json:"username" orm:"column(username);unique;size(100);"`
	Password string    `json:"-" orm:"column(password);size(300)"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

// 多字段索引
func (u *Group) TableIndex() [][]string {
	return [][]string{
		{"Id", "Username"},
	}
}

// 多字段唯一键
func (u *Group) TableUnique() [][]string {
	return [][]string{
		{"Username", "Email"},
	}
}

func (u *Group) TableName() string {
	return TableName("user")
}

func init() {
	orm.RegisterModel(new(Group))
}

func Groups() orm.QuerySeter {
	return OrmManager().QueryTable(new(Group))
}
