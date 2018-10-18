package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
)

var CachePool *redis.Pool

func init() {
	CachePool = NewRedisPool(beego.AppConfig.String("cache_redis"), beego.AppConfig.String("cache_redis"), 1)
}
