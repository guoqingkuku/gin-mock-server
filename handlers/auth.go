package handlers

import (
	"net/http"

	"gin-mock-server/models"

	"github.com/gin-gonic/gin"
)

// auth.go 认证相关的 Handler 示例

// LoginHandler 模拟登录接口 密码登录
func LoginHandler(c *gin.Context) {
	// 这里直接返回一个模拟的登录成功结果，实际项目中需要验证用户名和密码
	// 读取请求体中的用户名和密码（可选）
	var req struct {
		Username string `json:"userName" binding:"required"`
		Password string `json:"password" binding:"required"`
		Captcha  string `json:"captcha" binding:"omitempty"` // 可选字段，登录时可能需要验证码
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResult{
			Code: 400,
			Msg:  "参数校验失败: " + err.Error(),
			Data: nil,
		})
		return
	}

	// 模拟登录成功，返回一个假的 JWT token

	c.JSON(http.StatusOK, models.APIResult{
		Code: 0,
		Msg:  "登录成功",
		Data: map[string]string{
			"token": "mocked-jwt-token",
		},
	})
}

// 退出登录接口（可选）
func LogoutHandler(c *gin.Context) {
	// 模拟退出登录成功
	c.JSON(http.StatusOK, models.APIResult{
		Code: 0,
		Msg:  "退出登录成功",
		Data: nil,
	})
}
