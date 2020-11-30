package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/cached"
	"github.com/linqiurong2021/gin-book-backend/models"
	"github.com/linqiurong2021/gin-book-backend/services"
	"github.com/linqiurong2021/gin-book-frontend/dao"
	logic "github.com/linqiurong2021/gin-book-frontend/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"
)

// Login 登录
func Login(c *gin.Context) (string, bool) {
	//
	// var login = new(dao.Login) // binding 校验无效
	var login dao.Login       // binding 校验有效
	err := c.BindJSON(&login) // 绑定并校验
	// 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return "", false
	}
	// 验证码校验
	if !validator.CheckCode(login.Code) {
		c.JSON(http.StatusBadRequest, utils.BadRequest("code invalidate", ""))
		return "", false
	}
	admin, err := services.GetAdminByNameAndEncryptPassword(login.UserName, logic.MD5Encrypt(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return "", false
	}
	if admin == nil {
		c.JSON(http.StatusForbidden, utils.Forbidden("admin name or password invalidate", ""))
		return "", false
	}
	//
	singString, err := JWTToken(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return "", false
	}

	return singString, true
}

// JWTToken  JSONWebToken
func JWTToken(admin *models.Admin) (string, error) {
	//
	return cached.Create(admin)
}

// NameExists 校验用户名是否存在
func NameExists(c *gin.Context, admin *models.Admin, ID uint) (exists bool, err error) {
	// 判断名称是否已存在
	outAdmin, rowsAffected, err := services.GetAdminByName(admin.UserName)

	if err != nil {
		if rowsAffected != 0 {
			return false, err
		}
	}
	//
	if outAdmin != nil {
		//
		if outAdmin.ID == ID {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

// PhoneExists 校验手机号是否存在
func PhoneExists(c *gin.Context, admin *models.Admin, ID uint) (exists bool, err error) {
	// 判断手机号是否已存在
	outAdmin, rowsAffected, err := services.GetAdminByPhone(admin.Phone)
	// rowAffected = 0 说明是 查无数据
	if err != nil {
		if rowsAffected != 0 {
			return false, err
		}
	}
	//
	// 如果有传ID 则判断是不是当前自已
	if outAdmin != nil {
		// 如果为当前自已则返回
		if outAdmin.ID == ID {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

// CreateAdmin 创建用户
func CreateAdmin(c *gin.Context) (ok bool, err error) {
	var admin models.Admin
	err = c.ShouldBindJSON(&admin) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	// 判断名称是否存在
	exists, err := NameExists(c, &admin, cached.Admin.ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, &admin, cached.Admin.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return
	}
	//
	// 密码加密
	admin.Password = logic.MD5Encrypt(admin.Password)

	outAdmin, err := services.CreateAdmin(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("create success", outAdmin))
	//
	return true, nil

}

// NameAndPhoneExists 名称与手机号校验
func NameAndPhoneExists(c *gin.Context, admin *models.Admin, ID uint) (ok bool, err error) {
	// 判断名称是否存在
	exists, err := NameExists(c, admin, ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, admin, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return false, nil
	}
	return
}

// UpdateAdmin 更新用户
func UpdateAdmin(c *gin.Context) (ok bool, err error) {

	var admin models.Admin
	err = c.BindJSON(&admin) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	fmt.Println(cached.Admin.ID, "cached.Admin.ID")
	// 判断名称是否存在
	exists, err := NameExists(c, &admin, cached.Admin.ID)
	if err != nil {
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("user_name has exists", ""))
		return false, nil
	}
	// 判断手机号是否存在
	exists, err = PhoneExists(c, &admin, cached.Admin.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, err
	}
	if exists {
		c.JSON(http.StatusBadRequest, utils.BadRequest("phone has exists", ""))
		return false, nil
	}
	admin.ID = cached.Admin.ID
	//
	admin.Password = logic.MD5Encrypt(admin.Password)
	outAdmin, err := models.UpdateAdmin(&admin)
	if err != nil {
		return false, err
	}
	// 创建成功
	c.JSON(http.StatusOK, utils.Success("update success", outAdmin))
	// 获取当前
	return true, nil
}

// Logout 退出登录
func Logout(c *gin.Context) {
	// 直接清除
	// 数据清除
	// 退出成功
	c.JSON(http.StatusOK, utils.Success("logout success", ""))
}

// ListAdminByPage 用户列表分页
func ListAdminByPage(c *gin.Context) {
	//
	var page dao.Page
	c.BindQuery(&page)
	//
	list, total, err := services.GetListAdminByPage(page.Page, page.PageSize)
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
