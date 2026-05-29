package models

// User 用户数据模型
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Email string `json:"email,omitempty"` // 可选字段
	Nickname string `json:"nickname,omitempty"` // 可选字段
}

// CreateUserRequest 创建用户请求体（带校验）
type CreateUserRequest struct {
	Name string `json:"name" binding:"required,min=2"` // 必填，最少2个字符
	Age  int    `json:"age" binding:"required,gte=1"`  // 必填，必须大于等于1
	Email string `json:"email,omitempty" binding:"omitempty,email"` // 可选字段，必须是有效的邮箱格式
	Nickname string `json:"nickname,omitempty" binding:"omitempty,min=2,max=100"` // 可选字段，最少2个字符，最多100个字符
}