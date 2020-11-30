package services

import (
	"github.com/linqiurong2021/gin-book-backend/models"
)

// CreateAdmin 创建管理员
func CreateAdmin(inAdmin *models.Admin) (outAdmin *models.Admin, err error) {
	return models.CreateAdmin(inAdmin)
}

// GetAdminByID 通过管理员ID获取
func GetAdminByID(adminID uint) (admin *models.Admin, err error) {
	return models.GetAdminByID(adminID)
}

// GetAdminByPhone 通过手机号获取管理员信息
func GetAdminByPhone(phone string) (admin *models.Admin, rowsAffected int64, err error) {
	return models.GetAdminByPhone(phone)
}

// GetAdminByName 通过名称获取管理员信息
func GetAdminByName(name string) (admin *models.Admin, rowsAffected int64, err error) {
	return models.GetAdminByName(name)
}

// GetAdminByFieldValue 通过名称获取管理员信息
func GetAdminByFieldValue(field string, value string) (admin *models.Admin, rowsAffected int64, err error) {
	return models.GetAdminByFieldValue(field, value)
}

// GetAdminByNameAndEncryptPassword 通过管理员名或密码
func GetAdminByNameAndEncryptPassword(adminName string, encryptPassword string) (admin *models.Admin, err error) {
	return models.GetAdminByNameAndEncryptPassword(adminName, encryptPassword)
}

// UpdateAdmin 更新数据
func UpdateAdmin(info *models.Admin) (admin *models.Admin, err error) {

	return models.UpdateAdmin(info)
}

// DeleteAdminByID 通过ID删除管理员
func DeleteAdminByID(adminID int) (admin *models.Admin, err error) {

	return models.DeleteAdminByID(adminID)
}

// GetListAdminByPage 管理员列表分页
func GetListAdminByPage(page int, pageSize int) (adminList []*models.Admin, count int64, err error) {
	adminList, count, err = models.GetListAdminByPage(page, pageSize)
	return
}
