package redisp

import (
	"github.com/astaxie/beego"
)

var (
	db, _     = beego.AppConfig.Int("cache_redis_db")
	CachePool = NewRedisPool(beego.AppConfig.String("cache_redis_url"), beego.AppConfig.String("cache_redis_password"), db) // cache_redis
)
