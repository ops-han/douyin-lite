package middleware

import (
	"douyin-lite/comm"
	"douyin-lite/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg": "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := comm.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg": "权限不足",
			})
			c.Abort()
			return
		}
		userId := claims.UserId
		DB := comm.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg": "权限不足",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}