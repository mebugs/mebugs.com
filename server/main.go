package main

import (
	"log"
	"net/http"
	"server/config"
	"server/router"
)

// 主函数
func main() {
	log.Printf("Listening on port %s \n", config.RunConfig.Server.Port)
	// 启用HTTP服务
	err := http.ListenAndServe(config.RunConfig.Server.Port, router.Routers)
	if err != nil {
		log.Fatal(err)
	}
}
