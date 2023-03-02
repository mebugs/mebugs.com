package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"siteOl.com/stone/server/src/utils/logUtil"
)

var RunConfig *Config

// Config 配置对象
type Config struct {
	Server Server // 服务配置
	MySQL  MySQL  // MySQL数据配置
}

// Server 服务配置
type Server struct {
	Port     string // 服务端口
	LogLevel string // 日志等级
	LogRoot  int    // 日志展示路径范围
}

// MySQL MySQL数据配置组
type MySQL struct {
	Platform string // 授权平台
	Blog     string // 博客应用
}

// 配置初始化
func init() {
	SetEnv("server")                             // 初始化环境变量
	ReadConfig()                                 // 读取
	logUtil.SetLevel(RunConfig.Server.LogLevel)  // 日志等级
	logUtil.SetAutoSplit(logUtil.CronDaily)      // 日志切割
	logUtil.SetAppRoot(RunConfig.Server.LogRoot) // 日志深度(2表示上级../logs/）
}

// ReadConfig 启动读取配置文件
func ReadConfig() {
	RunConfig = &Config{}
	filePath := fmt.Sprintf("%s/config/config_%s.json", WorkDir, SysEnv)
	file, _ := os.Open(filePath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	for decoder.More() {
		err := decoder.Decode(RunConfig)
		if err != nil {
			log.Printf("Read %s Error: %v \n", filePath, err)
		}
	}
	log.Printf("Config Read %s Success .\n", filePath)
}
