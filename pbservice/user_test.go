package pbservice

import (
	_ "github.com/zhwei820/appconfig/utils/gotests"

	"testing"
	_ "github.com/zhwei820/appconfig/routers"
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"reflect"
	. "github.com/zhwei820/appconfig/pb/appconfig"
	"fmt"
)

func TestSay(t *testing.T) {
	go func() {
		beego.Run()
	}()

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	client := rpc.NewHTTPClient("http://127.0.0.1:8081/rpc")
	var helloService *HelloService
	client.UseService(&helloService)

	b, _ := prepareArgs().Marshal()
	r, _ := helloService.Hello(b)

	reply := &SayOutput{}
	reply.Unmarshal(r)

	fmt.Printf("%v", reply.Title)
}

type HelloService struct {
	Hello func(args []byte) (reply []byte, err error) `simple:"true"`
}

func prepareArgs() *SayInput {
	b := true
	var i int32 = 100000
	var s = "许多往事在眼前一幕一幕，变的那麼模糊"

	var args SayInput

	v := reflect.ValueOf(&args).Elem()
	num := v.NumField()
	for k := 0; k < num; k++ {
		field := v.Field(k)
		if field.Type().Kind() == reflect.Ptr {
			switch v.Field(k).Type().Elem().Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.Set(reflect.ValueOf(&i))
			case reflect.Bool:
				field.Set(reflect.ValueOf(&b))
			case reflect.String:
				field.Set(reflect.ValueOf(&s))
			}
		} else {
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.SetInt(100000)
			case reflect.Bool:
				field.SetBool(true)
			case reflect.String:
				field.SetString(s)
			}
		}

	}
	return &args
}
