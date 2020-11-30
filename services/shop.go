package services

import "github.com/linqiurong2021/gin-book-backend/models"

// CreateShop 创建店铺
func CreateShop(inShop *models.Shop) (outShop *models.Shop, err error) {
	return models.CreateShop(inShop)
}

// GetShopByID 通过店铺ID获取
func GetShopByID(shopID uint) (shop *models.Shop, err error) {
	return models.GetShopByID(shopID)
}

// GetShopByName 通过名称获取店铺信息
func GetShopByName(name string) (shop *models.Shop, rowsAffected int64, err error) {
	return models.GetShopByName(name)
}

// GetShopByFieldValue 通过名称获取店铺信息
func GetShopByFieldValue(field string, value string) (shop *models.Shop, rowsAffected int64, err error) {
	return models.GetShopByFieldValue(field, value)
}

// UpdateShop 更新数据
func UpdateShop(info *models.Shop) (shop *models.Shop, err error) {

	return models.UpdateShop(info)
}

// DeleteShopByID 通过ID删除店铺
func DeleteShopByID(shopID int) (shop *models.Shop, err error) {

	return models.DeleteShopByID(shopID)
}

// GetListShopByPage 店铺列表分页
func GetListShopByPage(page int, pageSize int) (shopList []*models.Shop, count int64, err error) {
	shopList, count, err = models.GetListShopByPage(page, pageSize)
	return
}
