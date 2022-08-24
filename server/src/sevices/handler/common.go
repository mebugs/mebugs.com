package handler

import "net/http"

// 控制器对象
type Func func(http.ResponseWriter, *http.Request)

func (f Func) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
