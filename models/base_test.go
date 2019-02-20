package models

import (
	_ "github.com/zhwei820/natsmicro/utils/gotests"

	"testing"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"

	"github.com/zhwei820/natsmicro/utils/random"
)

func TestTransaction(t *testing.T) {
	Init()
	testRollback(1)
	testRollback(0)
}

func testRollback(kk int)  {
	group:=Group{Group:"x",GroupName:"x",Permission:"x"}
	sqlFunc:= func(ormer orm.Ormer) error {
		_, err:=ormer.Insert(&group)
		if err!=nil{
			return err
		}
		group.Permission = "yyyy" + random.RandNumString(6)
		_, err=ormer.Update(&group, "Permission")
		if err!=nil{
			return err
		}
		if kk == 0{
			return errors.New("testtest")
		}else{
			return nil
		}
	}
	Transaction(sqlFunc)
}