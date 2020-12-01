package logic

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/models"
	"github.com/linqiurong2021/gin-book-backend/services"
	"github.com/linqiurong2021/gin-book-frontend/dao"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"
	"gorm.io/gorm"
)

// CreateCate 创建分类
func CreateCate(c *gin.Context) (ok bool, err error) {
	var cate models.Cate
	err = c.ShouldBindJSON(&cate) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}

	outCate, err := services.CreateCate(&cate)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("create success", outCate))
	//
	return true, nil

}

// UpdateCate 更新分类
func UpdateCate(c *gin.Context) (ok bool, err error) {

	var cate models.Cate
	err = c.BindJSON(&cate) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}

	outCate, err := models.UpdateCate(&cate)
	if err != nil {
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("update success", outCate))
	// 获取当前
	return true, nil
}

// GetListCateTree 获取列表 不分页  (暂用不到)
func GetListCateTree(c *gin.Context) (cateList []*models.Cate, err error) {
	cateList, err = services.GetListCateTree()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
	}
	// listTree := services.ConverListToTree(cateList)
	c.JSON(http.StatusOK, utils.Success("get success", cateList))
	return
}

// DeleteCate 删除
func DeleteCate(c *gin.Context) (ok bool, err error) {
	var delete dao.ID
	c.BindUri(&delete)
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	//
	_, err = services.DeleteCateByID(delete.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success("delete success", ""))
	return true, nil
}

// ConverListToTree 转树形结构
func ConverListToTree() {
	// services.ConverListToTree()
	return
}
