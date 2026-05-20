package main

import (
	"gin-mock-server/handlers"
	"gin-mock-server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建默认路由引擎（自带 Logger 和 Recovery 崩溃恢复中间件）
	r := gin.Default()

	// 2. 挂载自定义跨域中间件
	r.Use(middleware.Cors())

	// 3. 路由分组（版本控制，方便以后扩展 v2）
	v1 := r.Group("/api/v1")
	{
		// GET 请求示例：http://localhost:8080/api/v1/users/123
		v1.GET("/users/:id", handlers.GetUserHandler)
		
		// POST 请求示例：http://localhost:8080/api/v1/users
		v1.POST("/users", handlers.CreateUserHandler)
	}

	// 4. 启动服务，监听 8080 端口
	r.Run(":8080")
}