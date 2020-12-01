package models

import (
	"github.com/linqiurong2021/gin-book-frontend/mysql"
)

// Cate 分类
type Cate struct {
	Common   `gorm:"embedded"`
	Name     string `json:"name" gorm:"name;type:varchar(20);" binding:"required,min=2,max=20" label:"分类名" `
	Order    uint   `json:"order" gorm:"order" label:"排序"`
	Note     string `json:"note" gorm:"note;type:varchar(100)" binding:"required,max=100" label:"备注"`
	ParentID uint   `json:"parent_id" gorm:"parent_id;type:bigint;default=0;" label:"上级ID"`
	Cate     []Cate `json:"children" gorm:"foreignKey:ParentID;references:ID" `
}

// CreateCate 创建分类
func CreateCate(inCate *Cate) (cate *Cate, err error) {

	if err := mysql.DB.Create(&inCate).Error; err != nil {
		return nil, err
	}
	cate = inCate
	return
}

// GetCateByID 通过ID获取分类信息
func GetCateByID(cateID uint) (cate *Cate, err error) {
	if err := mysql.DB.Where("id = ?", cateID).Find(&cate).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateCate 更新数据
func UpdateCate(info *Cate) (outCate *Cate, err error) {
	if err := mysql.DB.Where("id = ?", info.ID).Save(info).Error; err != nil {
		return nil, err
	}
	outCate = info
	return
}

// DeleteCateByID 通过ID删除分类
func DeleteCateByID(cateID int) (cate *Cate, err error) {
	if err := mysql.DB.Where("id = ?", cateID).Delete(&Cate{}).Error; err != nil {
		return nil, err
	}
	return
}

// GetListCateTree 获取列表树
func GetListCateTree() (cateList []*Cate, err error) {
	if err := mysql.DB.Find(&cateList).Error; err != nil {
		return nil, err
	}

	return cateList, nil
}
