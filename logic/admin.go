package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// Login 登录
func Login(c *gin.Context) (token string, ok bool) {
	singString := "singString"
	c.JSON(http.StatusOK, utils.Success("login success", singString))
	return "", true
}

// Logout 退出登录
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("logout success", ""))
	return
}
