package utils

import (
	"time"
	"github.com/zhwei820/appconfig/utils/define"
)

func GetExpireTime() int64 {
	return time.Now().Unix() + int64(define.CacheDuration)
}
