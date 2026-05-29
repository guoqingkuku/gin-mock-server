
package models

// book.go 书籍数据模型
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description,omitempty"` // 可选字段
}