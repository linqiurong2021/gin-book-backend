package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// Login 登录
func Login(c *gin.Context) {
	singString, ok := logic.Login(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, utils.Success("login success", singString))
}

// Logout 退出登录
func Logout(c *gin.Context) {
	logic.Logout(c)
}

// CreateAdmin 创建管理员
func CreateAdmin(c *gin.Context) {

}

// UpdateAdmin 更新管理员
func UpdateAdmin(c *gin.Context) {

}

// DeleteAdmin 删除管理员
func DeleteAdmin(c *gin.Context) {

}

// ListAdminByPage 管理员分页
func ListAdminByPage(c *gin.Context) {

}
