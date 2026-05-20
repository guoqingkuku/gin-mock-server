package models

// User 用户数据模型
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUserRequest 创建用户请求体（带校验）
type CreateUserRequest struct {
	Name string `json:"name" binding:"required,min=2"` // 必填，最少2个字符
	Age  int    `json:"age" binding:"required,gte=1"`  // 必填，必须大于等于1
}