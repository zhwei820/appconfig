package pbservice

import (
	"github.com/zhwei820/natsmicro/pbservice/say"
	"github.com/zhwei820/natsmicro/pbservice/sing"
)

var ServiceConfig = map[string]func([]byte) ([]byte, error){
	"sayhello":  say.SayHello,
	"singhello": sing.SingHello,
}
