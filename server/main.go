package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"siteOl.com/stone/server/src/data/mysql"

	"siteOl.com/stone/server/src/utils/commUtil"
	"siteOl.com/stone/server/src/utils/logUtil"

	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/router"
)

// 主函数
func main() {
	// 初始化数据库
	mysql.InitMySQL()
	// 初始化路由
	router := router.NewRouter()
	httpServer := &http.Server{Addr: config.RunConfig.Server.Port, Handler: router}
	// 启用HTTP服务 - 注册自定义路由
	go commUtil.RecoverWrap(func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logUtil.ErrorF("Server Listening on port %s . Err %v", config.RunConfig.Server.Port, err)
			os.Exit(1)
		}
	})
	logUtil.PrintF("Server Listening on port %s", config.RunConfig.Server.Port)
	// 优雅关机
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logUtil.PrintF("Server Get a signal %s, Stop the consume process", sig.String())
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			// gracefully shutdown with timeout
			_ = httpServer.Shutdown(ctx)
			return
		}
	}
}
