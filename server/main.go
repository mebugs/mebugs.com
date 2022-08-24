package main

import (
	"log"
	"net/http"
	"server/config"
	"server/src/router"
)

// 主函数
func main() {
	log.Printf("Listening on port %s \n", config.RunConfig.Server.Port)
	// 启用HTTP服务 - 注册自定义路由
	err := http.ListenAndServe(config.RunConfig.Server.Port, router.Routers)
	if err != nil {
		log.Fatal(err)
	}
}
