package config

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strings"
)

var RunConfig *Config

// 配置对象
type Config struct {
	Server Server // 服务配置
	MySQL  MySQL  // MySQL数据配置
}

// 服务配置
type Server struct {
	Port string // 服务端口
}

// MySQL数据配置组
type MySQL struct {
	Platform string // 授权平台
	Blog     string // 博客应用
}

// 配置初始化
func init() {
	ReadConfig()
}

// 启动读取配置文件
func ReadConfig() {
	RunConfig = &Config{}
	wordDir, _ := os.Getwd()
	if runtime.GOOS == "windows" {
		wordDir = strings.ReplaceAll(wordDir, "\\", "/")
	}
	file, _ := os.Open(wordDir + "/config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	for decoder.More() {
		err := decoder.Decode(RunConfig)
		if err != nil {
			log.Printf("Error: %v \n", err)
		}
	}
	log.Printf("Config Read Success .\n")
}
