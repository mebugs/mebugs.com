package utils

import (
	"math/rand"
	"strings"
	"time"
)

// ArrayInclude 字符串是否位于数组中
func ArrayInclude(node string, nodes []string) bool {
	if len(nodes) == 0 {
		return false
	}
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

// ArrayStartWith 字符串是否以某开头
func ArrayStartWith(node string, nodes []string) bool {
	if len(nodes) == 0 {
		return false
	}
	for _, n := range nodes {
		if strings.HasPrefix(node, n) {
			return true
		}
	}
	return false
}

// ShuffleAndSelect 洗牌算法，指定位数的新数组
func ShuffleAndSelect(arr []string, count int) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr[:count]
}
