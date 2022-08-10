package router

import (
	"fmt"
	"net/http"
	"regexp"
)

// 路由对象
type Router struct {
	routers       map[string]http.Handler // 路由列表
	middlewareFun []MDFunc                // 中间件列表
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	find := false
	// 遍历路由开始匹配
	for route, handler := range r.routers {
		matched, err := regexp.MatchString(route, req.URL.Path)
		if err != nil {
			fmt.Errorf("Match Router Fail  %s", err.Error())
			http.Error(w, "Router Err", 404)
			return
		}
		// 未匹配到数据
		if !matched {
			continue
		}
		// 读取路由方法
		find = true
		handler.ServeHTTP(w, req)
		break
	}
	if !find {
		fmt.Errorf("Match Router Never Find  %s", req.URL.Path)
		http.Error(w, "Router Err", 404)
	}
}

// 应用中间件
func (r *Router) Use(m MDFunc) {
	r.middlewareFun = append(r.middlewareFun, m)
}

// 注册路由
func (r *Router) Register(path string, handler http.Handler) {
	// 组合handler = handler + 中间件
	var mergedHandler = handler
	if len(r.middlewareFun) > 0 {
		// 依照正常的中间件顺序写入（先写的在前，后写的在后）
		for _, mdFunc := range r.middlewareFun {
			mergedHandler = mdFunc(mergedHandler)
		}
	}
	// 最终组合的 handler
	r.routers[path] = mergedHandler
}
