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

// shuffleAndSelect 洗牌算法，指定位数的新数组
func shuffleAndSelect(arr []string, count int) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr[:count]
}

// ShuffleAndSelectEx 洗牌算法，指定位数的新数组，排除部分数据
func ShuffleAndSelectEx(arr, extArr []string, count int) []string {
	extMap := map[string]bool{}
	for _, ext := range extArr {
		extMap[ext] = true
	}
	newArr := make([]string, 0)
	for _, a := range arr {
		if !extMap[a] {
			newArr = append(newArr, a)
		}
	}
	if len(newArr) <= count {
		extArr = append(extArr, newArr...)
		return extArr
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(newArr), func(i, j int) {
		newArr[i], newArr[j] = newArr[j], newArr[i]
	})
	extArr = append(extArr, newArr[:count]...)
	return extArr
}

// ArrayToSet 字符串数组去重
func ArrayToSet(arr []string) []string {
	outArr := make([]string, 0)
	arrMap := map[string]bool{}
	for _, ext := range arr {
		if !arrMap[ext] {
			outArr = append(outArr, ext)
			arrMap[ext] = true
		}
	}
	return outArr
}

// GetFirstAndExt 获取前面的数据并忽略存在的
func GetFirstAndExt(arr, extArr []string, count int) []string {
	extMap := map[string]bool{}
	for _, ext := range extArr {
		extMap[ext] = true
	}
	for _, a := range arr {
		if count == 0 {
			return extArr
		}
		if !extMap[a] {
			extArr = append(extArr, a)
			extMap[a] = true
			count--
		}
	}
	return extArr
}
