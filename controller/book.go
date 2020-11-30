package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/logic"
)

// ListBookByPage 书籍分页
func ListBookByPage(c *gin.Context) {
	logic.ListBookByPage(c)
	return
}
