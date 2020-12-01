package services

import "github.com/linqiurong2021/gin-book-backend/models"

// CreateCate 创建分类
func CreateCate(inCate *models.Cate) (outCate *models.Cate, err error) {
	return models.CreateCate(inCate)
}

// GetCateByID 通过分类ID获取
func GetCateByID(cateID uint) (cate *models.Cate, err error) {
	return models.GetCateByID(cateID)
}

// UpdateCate 更新数据
func UpdateCate(info *models.Cate) (cate *models.Cate, err error) {

	return models.UpdateCate(info)
}

// DeleteCateByID 通过ID删除分类
func DeleteCateByID(cateID int) (cate *models.Cate, err error) {

	return models.DeleteCateByID(cateID)
}

// GetListCateTree 分类树
func GetListCateTree() (cateList []*models.Cate, err error) {
	cateList, err = models.GetListCateTree()
	return
}

// ConverListToTree 列表转树形结构
func ConverListToTree(cateList []*models.Cate) {
	//
	return
}
