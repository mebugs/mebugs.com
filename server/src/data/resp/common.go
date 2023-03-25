package resp

import (
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

// 定义一些常量
var (
	// OK Json默认成功返回
	OK = Success(nil)
	// SysErr Json默认系统异常
	SysErr = Error()
)

type ResBody struct {
	Code int         `json:"code"` // 响应码
	Msg  string      `json:"msg"'` // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// Success Json数据返回
func Success(data interface{}) ResBody {
	return jsonResult(http.StatusOK, "", data)
}

// Validate Json校验返回400（已翻译）
func Validate(err error) ResBody {
	return jsonResult(http.StatusInternalServerError, err.Error(), nil)
}

// Error Json错误返回500
func Error() ResBody {
	return jsonResult(http.StatusInternalServerError, "", nil)
}

// 公共调用
func jsonResult(code int, msg string, data interface{}) ResBody {
	resp := ResBody{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return resp
}
