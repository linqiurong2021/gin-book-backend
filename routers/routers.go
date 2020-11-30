package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/controller"
	"github.com/linqiurong2021/gin-book-backend/middlewares"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 需要校验的分组
		authGroup(v1)
		// 不需要校验
		noAuthGroup(v1)
		// 默认路由
		defaultRouter(r)
	}
}

func notMethod(c *gin.Context) {
	// utils
	c.JSON(http.StatusBadRequest, utils.BadRequest("bad request", ""))
	c.Abort()
}

// 未匹配到路由时
func noRoute(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.BadRequest("not found route", ""))
	c.Abort()
}

// 系统默认的路由
func defaultRouter(r *gin.Engine) {

	r.NoMethod(notMethod)
	// 未匹配到路由时
	r.NoRoute(noRoute)
	// 心跳检测
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.Success("pong", ""))
	})

}

// noAuthGroup 不需要登录校验
func noAuthGroup(version *gin.RouterGroup) {
	version.POST("/login", controller.Login)
	version.GET("/logout", controller.Logout)
}

func authGroup(version *gin.RouterGroup) {
	authGroup := version.Group("")
	authGroup.Use(middlewares.JWTAdminTokenCheck())
	{
		userGroup(authGroup)
		shopGroup(authGroup)
		cateGroup(authGroup)
		adminGroup(authGroup)
		shopUserGroup(authGroup)
	}
}

// userGroup 用户分组路由
func userGroup(version *gin.RouterGroup) {
	user := version.Group("/user")
	user.GET("", controller.ListUserByPage)
}

// shopGroup 店铺分组路由
func shopGroup(version *gin.RouterGroup) {
	shop := version.Group("/shop")
	shop.GET("", controller.ListShopByPage)
	shop.POST("", controller.CreateShop)
	shop.PUT("", controller.UpdateShop)
	shop.DELETE("", controller.DeleteShop)
}

// shopUserGroup 分类分组路由
func shopUserGroup(version *gin.RouterGroup) {
	shopUser := version.Group("/shop_user")
	shopUser.GET("", controller.ListShopUserByPage)
	shopUser.POST("", controller.CreateShopUser)
	shopUser.PUT("", controller.UpdateShopUser)
	shopUser.DELETE("", controller.DeleteShopUser)
}

// cateGroup 分类分组路由
func cateGroup(version *gin.RouterGroup) {
	cate := version.Group("/cate")
	cate.GET("", controller.ListCateByPage)
	cate.POST("", controller.CreateCate)
	cate.PUT("", controller.UpdateCate)
	cate.DELETE("", controller.DeleteCate)
}

// adminGroup 管理员分组路由
func adminGroup(version *gin.RouterGroup) {
	admin := version.Group("/admin")
	admin.GET("", controller.ListAdminByPage)
	admin.POST("", controller.CreateAdmin)
	admin.PUT("", controller.UpdateAdmin)
	admin.DELETE("", controller.DeleteMenu)
}
