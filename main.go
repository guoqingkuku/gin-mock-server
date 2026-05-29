package main

import (
	"gin-mock-server/handlers"
	"gin-mock-server/middleware"
	"gin-mock-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建默认路由引擎（自带 Logger 和 Recovery 崩溃恢复中间件）
	r := gin.Default()
	// 统一处理 404 路由找不到的情况
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.APIResult{
			Code: 404,
			Msg:  "接口不存在",
			Data: nil,
		})
	})

	// 2. 挂载自定义跨域中间件
	r.Use(middleware.Cors())

	// 3. 路由分组（版本控制，方便以后扩展 v2）
	v1 := r.Group("/api/v1")
	{
		// GET 请求示例：http://localhost:8080/api/v1/users/123
		v1.GET("/users/:id", handlers.GetUserHandler)

		// POST 请求示例：http://localhost:8080/api/v1/users
		v1.POST("/users", handlers.CreateUserHandler)
		// 书籍列表接口
		v1.GET("/books", handlers.ListBooksHandler)

		// 认证相关接口
		// 密码登录
		v1.POST("/auth/login", handlers.LoginHandler)
		// 退出登录
		v1.GET("/auth/logout", handlers.LogoutHandler)
	}

	// 4. 启动服务，监听 8080 端口
	r.Run(":8080")
}
