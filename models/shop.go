package models

import (
	"errors"
	"fmt"

	"github.com/linqiurong2021/gin-book-frontend/mysql"
	"gorm.io/gorm"
)

// Shop 店铺
type Shop struct {
	Common   `gorm:"embedded"`
	Name     string    `json:"name" gorm:"name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Contact  string    `json:"contact" gorm:"contact;type:char(11);" binding:"len=11" label:"手机号"`
	ShopUser *ShopUser `json:"shop_user"`
	Books    []*Book   `json:"books"`
}

// CreateShop 创建店铺
func CreateShop(inShop *Shop) (shop *Shop, err error) {

	if err := mysql.DB.Create(&inShop).Error; err != nil {
		return nil, err
	}
	shop = inShop
	return
}

// GetShopByID 通过ID获取店铺信息
func GetShopByID(shopID uint) (shop *Shop, err error) {
	if err := mysql.DB.Where("id = ?", shopID).Find(&shop).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateShop 更新数据
func UpdateShop(info *Shop) (outShop *Shop, err error) {
	if err := mysql.DB.Where("id = ?", info.ID).Save(info).Error; err != nil {
		return nil, err
	}
	outShop = info
	return
}

// DeleteShopByID 通过ID删除店铺
func DeleteShopByID(shopID int) (shop *Shop, err error) {
	if err := mysql.DB.Where("id = ?", shopID).Delete(&Shop{}).Error; err != nil {
		return nil, err
	}
	return
}

// GetShopByName 通过某个字段获取店铺信息
func GetShopByName(value string) (outShop *Shop, rowsAffected int64, err error) {
	return GetShopByFieldValue("user_name", value)
}

// GetShopByFieldValue 通过某个字段获取店铺信息
func GetShopByFieldValue(field string, value string) (outShop *Shop, rowsAffected int64, err error) {
	var shop = new(Shop)
	var where string = fmt.Sprintf("%s = ?", field)
	//
	record := mysql.DB.Where(where, value).First(&shop)
	// 查不到数据
	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, 0, record.Error
	}
	// 异常
	if record.Error != nil {
		return nil, 1, record.Error
	}

	return shop, 1, nil
}

// GetListShopByPage 获取列表 分页 (暂用不到)
func GetListShopByPage(page int, pageSize int) (shopList []*Shop, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&shopList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&Shop{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return shopList, count, nil
}
