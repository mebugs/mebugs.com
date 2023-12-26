package model

// CacheRouter 缓存路由对象
type CacheRouter struct {
	Type       uint8    `json:"type"`       // 免授权路由 0 授权 1 免授权
	ReqUnPrint bool     `json:"reqUnPrint"` // 请求日志不打印
	ReqSecure  []string `json:"reqSecure"`  // 请求日志脱敏数组字段
	ResUnPrint bool     `json:"resUnPrint"` // 响应日志不打印
	ResSecure  []string `json:"resSecure"`  // 响应日志脱敏数组字段
}
