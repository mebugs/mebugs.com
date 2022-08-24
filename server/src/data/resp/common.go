package resp

import (
	"encoding/json"
	"net/http"
)

/**
 *
 * 统一数据JSON返回结构
 *
 *
 * @author 米虫@mebugs.com
 * @since 2022-08-16
 */

// 404 路由获取失败
func RouterErr(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// 403 鉴权失败
func AuthErr(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

type ResBody struct {
	Code int         `json:"code"` // 响应码
	Msg  string      `json:"msg"'` // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// Json默认成功返回
func JsonOK(w http.ResponseWriter) {
	JsonSuccess(w, nil)
}

// Json数据返回
func JsonSuccess(w http.ResponseWriter, data interface{}) {
	jsonResult(w, http.StatusOK, http.StatusText(http.StatusOK), data)
}

// Json失败返回
func JsonFail(w http.ResponseWriter, msg string) {
	jsonResult(w, http.StatusInternalServerError, msg, nil)
}

// Json错误返回
func JsonError(w http.ResponseWriter, err error) {
	jsonResult(w, http.StatusInternalServerError, err.Error(), nil)
}

// 公共调用
func jsonResult(w http.ResponseWriter, code int, msg string, data interface{}) {
	resp := ResBody{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	respBytes, _ := json.Marshal(resp)
	// 200 返回body
	w.Write(respBytes)
}
