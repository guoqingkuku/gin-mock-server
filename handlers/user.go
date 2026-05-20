package handlers

import (
	"net/http"
	"strconv"

	"gin-mock-server/models"
	"github.com/gin-gonic/gin"
)

// GetUserHandler GET /api/v1/users/:id (获取单个用户)
func GetUserHandler(c *gin.Context) {
	// 获取路由参数 :id
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// 统一的 JSON 响应格式
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": models.User{
			ID:   id,
			Name: "Flutter Tester",
			Age:  18,
		},
	})
}

// CreateUserHandler POST /api/v1/users (创建用户，含 JSON 解析与参数校验)
func CreateUserHandler(c *gin.Context) {
	var req models.CreateUserRequest

	// ShouldBindJSON 会自动解析 POST 的 Body，并根据 models 里的 binding 标签校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数校验失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 模拟创建成功
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "user created",
		"data": models.User{
			ID:   999, // 模拟生成的ID
			Name: req.Name,
			Age:  req.Age,
		},
	})
}