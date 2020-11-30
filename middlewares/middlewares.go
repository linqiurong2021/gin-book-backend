package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/cached"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// JWTAdminTokenCheck JWTToken 校验
func JWTAdminTokenCheck() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, utils.BadRequest("token must", ""))
			c.Abort()
			return
		}
		jwtToken, err := cached.Parse(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			c.Abort()
			return
		}
		cliams, ok := cached.Check(jwtToken)
		fmt.Printf("%+v #### %s", cliams,"#####")
		if ok && jwtToken.Valid {
			// 存储当前用户信息
			cached.Save(cliams)
			// 可以存储到 c.Set("user",user)
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, cliams)
			c.Abort()
			return
		}
	}
}
