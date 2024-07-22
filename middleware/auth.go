package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"main/configs"
)

func GetKey(_ *jwt.Token)(interface{} , error){
	return  configs.MySecret, nil
}
type MyClaims struct{
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}
//生成token
func GenToken(userID int64) (string , error){
	claims := MyClaims{
		UserID : userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1*time.Hour)),
			Issuer: "jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	return token.SignedString(configs.MySecret)
}
//解析token
func ParseToken(tokenString string) (*MyClaims , error){
	var claims MyClaims
	token , err := jwt.ParseWithClaims(tokenString , &claims ,  GetKey)
	if err != nil{
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	if(claims.ExpiresAt.Before(time.Now())){
		token.Valid = false
	}
	if !token.Valid {
		fmt.Println("token is invalid")
		return nil , errors.New("token is invalid")
	}
	return &claims , nil
}
//token验证中间件
func VerifyToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = string(c.PostForm("token"))
			if token == "" {
				c.JSON(http.StatusNotFound, gin.H{"error":"token is empty"})
				c.Abort()
				return
			}
		}
		claims , err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"token is ERROR"})
			c.Abort()
			return
		}
		c.Set("userID" , claims.UserID)
		c.Next()
    } 
}

