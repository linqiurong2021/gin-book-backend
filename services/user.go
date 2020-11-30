package services

import "github.com/linqiurong2021/gin-book-backend/models"

// GetListUserByPage 通过ID删除用户
func GetListUserByPage(page int, pageSize int) (userList []*models.User, count int64, err error) {
	userList, count, err = models.GetListUserByPage(page, pageSize)
	return
}
