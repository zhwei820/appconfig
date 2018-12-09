package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Group struct {
	Id         int64     `json:"id" orm:"column(id);pk;auto;unique"`
	GroupName  string    `json:"groupname" orm:"column(groupname);unique;size(100)"`
	Group      string    `json:"group" orm:"column(group);unique;size(100)"`
	Permission string    `json:"permission" orm:"column(permission);size(500)"`
	Created    time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated    time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

// 多字段索引
func (u *Group) TableIndex() [][]string {
	return [][]string{
		{"GroupName"},
	}
}

// 多字段唯一键
func (u *Group) TableUnique() [][]string {
	return [][]string{
		{"GroupName",},
	}
}

func (u *Group) TableName() string {
	return TableName("group")
}

func init() {
	orm.RegisterModel(new(Group))
}

func Groups() orm.QuerySeter {
	return OrmManager().QueryTable(new(Group))
}
