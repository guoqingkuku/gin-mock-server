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
	c.JSON(http.StatusOK, models.APIResult{
		Code: 0,
		Msg:  "success",
		Data: models.User{
			ID:   id,
			Name: "Flutter Tester",
			Age:  18,
			Email: "flutter.tester@example.com",
			Nickname: "FlutterTester",
		},
	})
}

// CreateUserHandler POST /api/v1/users (创建用户，含 JSON 解析与参数校验)
func CreateUserHandler(c *gin.Context) {
	var req models.CreateUserRequest

	// ShouldBindJSON 会自动解析 POST 的 Body，并根据 models 里的 binding 标签校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResult{
			Code: 400,
			Msg:  "参数校验失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 模拟创建成功
	c.JSON(http.StatusCreated, models.APIResult{
		Code: 201,
		Msg:  "user created",
		Data: models.User{
			ID:   999, // 模拟生成的ID
			Name: req.Name,
			Age:  req.Age,
			Email: req.Email,
			Nickname: req.Nickname,
		},
	})
}
