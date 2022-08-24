package platHandler

import (
	"net/http"
	"server/src/data/resp"
)

// 登陆授权
// /auth/login
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	resp.JsonSuccess(w, nil)
}
