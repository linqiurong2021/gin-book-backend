package controller

import (
	"net/http"

	"github.com/linqiurong2021/gin-book-backend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"

	"github.com/gin-gonic/gin"
)

// CreateAdmin 新增用户
func CreateAdmin(c *gin.Context) {
	ok, err := logic.CreateAdmin(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusAccepted, utils.CreateFailure(err.Error(), ""))
			return
		}
		return
	}
}

// UpdateAdmin 更新
func UpdateAdmin(c *gin.Context) {
	ok, err := logic.UpdateAdmin(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteAdmin 删除
func DeleteAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Success("Delete success", ""))
}

// ListAdminByPage 列表分页
func ListAdminByPage(c *gin.Context) {
	logic.ListAdminByPage(c)
	return
}

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
