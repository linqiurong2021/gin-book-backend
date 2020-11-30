package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/models"
	"github.com/linqiurong2021/gin-book-backend/services"
	"github.com/linqiurong2021/gin-book-frontend/dao"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"
)

// CreateShop 创建管理员
func CreateShop(c *gin.Context) (ok bool, err error) {
	var shop models.Shop
	err = c.ShouldBindJSON(&shop) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	outShop, err := services.CreateShop(&shop)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("create success", outShop))
	//
	return true, nil

}

// UpdateShop 更新管理员
func UpdateShop(c *gin.Context) (ok bool, err error) {

	var shop models.Shop
	err = c.BindJSON(&shop) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}

	outShop, err := models.UpdateShop(&shop)
	if err != nil {
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("update success", outShop))
	// 获取当前
	return true, nil
}

// ListShopByPage 管理员列表分页
func ListShopByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total, err := services.GetListShopByPage(page.Page, page.PageSize)
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
