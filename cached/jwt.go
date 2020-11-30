package cached

import (
	"fmt"
	"time"

	"github.com/linqiurong2021/gin-book-backend/models"
	"github.com/linqiurong2021/gin-book-frontend/config"

	"github.com/dgrijalva/jwt-go"
)

// Claims 自定义
type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// Create 创建
func Create(admin *models.Admin) (string, error) {
	//
	mySigningKey := []byte(config.Conf.JWTSignKey)
	//
	now := time.Now().Unix()
	expiresAt := now + config.Conf.TokenExpireMinutes*60

	// fmt.Println(expiresAt, "expiresAt")
	claims := Claims{
		admin.ID,
		admin.UserName,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "gin-book",
			Subject:   "github.com/linqiurong2021",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singString, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", singString, err)
	return singString, err
}

// Parse 解析Token
func Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JWTSignKey), nil
	})
	return token, err
}

// Check Token校验
func Check(jwtToken *jwt.Token) (*Claims, bool) {
	claims, ok := jwtToken.Claims.(*Claims)
	return claims, ok
}
