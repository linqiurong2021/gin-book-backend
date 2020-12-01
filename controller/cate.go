package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// CreateCate 新增分类
func CreateCate(c *gin.Context) {

	ok, err := logic.CreateCate(c)
	if !ok {
		if err != nil {
			return
		}
		return
	}
}

// UpdateCate 更新
func UpdateCate(c *gin.Context) {
	ok, err := logic.UpdateCate(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
		return
	}
}

// DeleteCate 删除
func DeleteCate(c *gin.Context) {
	ok, err := logic.DeleteCate(c)
	if !ok {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return
		}
	}
}

// GetListCateTree 列表分页
func GetListCateTree(c *gin.Context) {
	logic.GetListCateTree(c)
	return
}
