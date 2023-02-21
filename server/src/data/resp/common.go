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

type ResBody struct {
	Code int         `json:"code"` // 响应码
	Msg  string      `json:"msg"'` // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// RouterErr 404 路由获取失败
func RouterErr() ResBody {
	return jsonResult(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
}

// AuthErr 403 鉴权失败
func AuthErr() ResBody {
	return jsonResult(http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
}

// JsonOK Json默认成功返回
func JsonOK() ResBody {
	return JsonSuccess(nil)
}

// JsonSuccess Json数据返回
func JsonSuccess(data interface{}) ResBody {
	return jsonResult(http.StatusOK, http.StatusText(http.StatusOK), data)
}

// JsonFail Json失败返回
func JsonFail(msg string) ResBody {
	return jsonResult(http.StatusInternalServerError, msg, nil)
}

// JsonError Json错误返回
func JsonError(err error) ResBody {
	return jsonResult(http.StatusInternalServerError, err.Error(), nil)
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
