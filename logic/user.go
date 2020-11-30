package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/services"
	"github.com/linqiurong2021/gin-book-frontend/dao"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// ListUserByPage 用户列表分页
func ListUserByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total, err := services.GetListUserByPage(page.Page, page.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	listPage := &dao.ListPage{
		Total: total,
		List:  list,
	}

	c.JSON(http.StatusOK, utils.Success("get success", listPage))
	return
}
