package util

import (
	"testing"
	"github.com/garyburd/redigo/redis"
	"github.com/henrylee2cn/goutil"
)

func TestNewPool(t *testing.T) {
	pool := NewRedisPool("127.0.0.1:6379","redis123", 0)
	c:=pool.Get()
	_, err:=c.Do("SET","TEST", 111)
	if err!=nil{
		Llog.Error(err)
	}
	res, err:=redis.Int(c.Do("GET","TEST"))
	if err!=nil{
		Llog.Error(err)
	}
	Llog.Info(res)
}

func BenchmarkNewRedisPool(t *testing.B) {

	r := goutil.NewRandom("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_", '-', '_')

	pool := NewRedisPool("127.0.0.1:6379","redis123", 0)

	for ii:=0;ii<t.N;ii++ {
		c := pool.Get()

		_, _ = c.Do("HSET", "hsettest",r.RandomString(10), 111)
		c.Close()  // important
	}
}