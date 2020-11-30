package models

import (
	"errors"
	"fmt"

	"github.com/linqiurong2021/gin-book-frontend/mysql"
	"gorm.io/gorm"
)

// Admin 用户
type Admin struct {
	Common   `gorm:"embedded"`
	UserName string `json:"user_name" gorm:"user_name;type:varchar(20);" binding:"required,min=6,max=20" label:"用户名" `
	Password string `json:"password" gorm:"password;type:varchar(32);" binding:"required,min=6,max=20" label:"密码"`
	Phone    string `json:"phone" gorm:"phone;type:char(11);" binding:"len=11" label:"手机号"`
}

// CreateAdmin 创建管理员
func CreateAdmin(inAdmin *Admin) (admin *Admin, err error) {

	if err := mysql.DB.Create(&inAdmin).Error; err != nil {
		return nil, err
	}
	admin = inAdmin
	return
}

// GetAdminByID 通过ID获取管理员信息
func GetAdminByID(adminID uint) (admin *Admin, err error) {
	if err := mysql.DB.Where("id = ?", adminID).Find(&admin).Error; err != nil {
		return nil, err
	}
	return
}

// GetAdminByNameAndEncryptPassword 通过管理员名或密码
func GetAdminByNameAndEncryptPassword(adminName string, encryptPassword string) (outAdmin *Admin, err error) {
	var admin = new(Admin)
	record := mysql.DB.Where("user_name = ?", adminName).Where("password = ?", encryptPassword).Find(&admin)
	// 为空或查找数量为0时
	if errors.Is(record.Error, gorm.ErrRecordNotFound) || record.RowsAffected == 0 {
		return nil, nil
	}
	if record.Error != nil {
		// fmt.Print("BBBBB###")
		return nil, record.Error
	}
	return admin, nil
}

// UpdateAdmin 更新数据
func UpdateAdmin(info *Admin) (outAdmin *Admin, err error) {
	if err := mysql.DB.Where("id = ?", info.ID).Save(info).Error; err != nil {
		return nil, err
	}
	outAdmin = info
	return
}

// DeleteAdminByID 通过ID删除管理员
func DeleteAdminByID(adminID int) (admin *Admin, err error) {
	if err := mysql.DB.Where("id = ?", adminID).Delete(&Admin{}).Error; err != nil {
		return nil, err
	}
	return
}

// GetAdminByPhone 通过某个字段获取管理员信息
func GetAdminByPhone(value string) (outAdmin *Admin, rowsAffected int64, err error) {
	return GetAdminByFieldValue("phone", value)
}

// GetAdminByName 通过某个字段获取管理员信息
func GetAdminByName(value string) (outAdmin *Admin, rowsAffected int64, err error) {
	return GetAdminByFieldValue("user_name", value)
}

// GetAdminByFieldValue 通过某个字段获取管理员信息
func GetAdminByFieldValue(field string, value string) (outAdmin *Admin, rowsAffected int64, err error) {
	var admin = new(Admin)
	var where string = fmt.Sprintf("%s = ?", field)
	//
	record := mysql.DB.Where(where, value).First(&admin)
	// 查不到数据
	if errors.Is(record.Error, gorm.ErrRecordNotFound) {
		return nil, 0, record.Error
	}
	// 异常
	if record.Error != nil {
		return nil, 1, record.Error
	}

	return admin, 1, nil
}

// GetListAdminByPage 获取列表 分页 (暂用不到)
func GetListAdminByPage(page int, pageSize int) (adminList []*Admin, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&adminList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&Admin{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return adminList, count, nil
}
