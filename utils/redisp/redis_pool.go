package redisp

import (
	"time"
	"github.com/gomodule/redigo/redis"
	"strings"
	"log"
)

func NewRedisPool(server, password string, redisDb int) *redis.Pool {
	if ! strings.Contains(server, ":") {
		log.Fatalf("redis server url error! %v", server)
	}
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" { // 密码
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", redisDb); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
