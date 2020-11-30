package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/logic"
)

// ListUserByPage 用户分页
func ListUserByPage(c *gin.Context) {
	logic.ListUserByPage(c)
	return
}
