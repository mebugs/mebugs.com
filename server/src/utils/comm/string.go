package comm

import (
	"math/rand"
	"time"
)

var baseStr = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"

// TraceID 生成8位随机日志ID
func TraceID() string {
	return RandStr(8)
}

// RandStr 生成指定位数的随机字符
func RandStr(length int) string {
	bytes := []byte(baseStr)
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(10000)))
	for i := 0; i < length; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(result)
}
