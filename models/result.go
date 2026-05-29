package models

// APIResult 统一的 API 响应格式
type APIResult struct {
	Code int         `json:"code"` // 状态码，0表示成功，非0表示错误
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 具体数据
}
