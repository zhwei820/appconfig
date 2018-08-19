package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

// 数据库连接初始化
func Init() {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		beego.AppConfig.String("db_user"),
		beego.AppConfig.String("db_passwd"),
		beego.AppConfig.String("db_host"),
		beego.AppConfig.String("db_port"),
		beego.AppConfig.String("db_name"),
		beego.AppConfig.String("db_charset"))
	orm.RegisterDataBase("default", "mysql", conn)

	//自动建表
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("db_prefix") + str
}

// orm管理器
func OrmManager() orm.Ormer {
	return orm.NewOrm()
}