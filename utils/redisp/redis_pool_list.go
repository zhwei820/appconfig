package redisp

import (
	"github.com/astaxie/beego"
)

var CachePool = NewRedisPool(beego.AppConfig.String("cache_redis"), beego.AppConfig.String("cache_redis"), 1)
