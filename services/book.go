package services

import "github.com/linqiurong2021/gin-book-backend/models"

// GetListBookByPage 书籍分页
func GetListBookByPage(page int, pageSize int) (bookList []*models.Book, count int64, err error) {
	bookList, count, err = models.GetListBookByPage(page, pageSize)
	return
}
