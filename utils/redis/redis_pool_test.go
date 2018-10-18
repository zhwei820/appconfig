package redis

import (
	"testing"
	"github.com/garyburd/redigo/redis"
	"github.com/henrylee2cn/goutil"
	"github.com/rs/zerolog/log"
)

func TestNewPool(t *testing.T) {
	pool := NewRedisPool("127.0.0.1:6379", "redis123", 0)
	c := pool.Get()
	_, err := c.Do("SET", "TEST", 111)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	res, err := redis.Int(c.Do("GET", "TEST"))
	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Error().Int("test", res)

}

func BenchmarkNewRedisPool(t *testing.B) {

	r := goutil.NewRandom("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

	pool := NewRedisPool("127.0.0.1:6379", "redis123", 0)

	for ii := 0; ii < t.N; ii++ {
		c := pool.Get()

		_, _ = c.Do("HSET", "hsettest", r.RandomString(10), 111)
		c.Close() // important
	}
}
