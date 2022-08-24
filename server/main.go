package main

import (
	"log"
	"net/http"
	config2 "server/src/data/config"
	"server/src/router"
)

// 主函数
func main() {
	log.Printf("Server Listening on port %s . \n", config2.RunConfig.Server.Port)
	// 启用HTTP服务 - 注册自定义路由
	err := http.ListenAndServe(config2.RunConfig.Server.Port, router.Routers)
	if err != nil {
		log.Fatal(err)
	}
}
