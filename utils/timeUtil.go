package utils

import (
	"time"
)

func GetExpireTime() int64 {
	return time.Now().Unix() + 3600 * 24 * 7
}