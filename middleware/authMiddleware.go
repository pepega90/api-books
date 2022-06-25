package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin_learn/models/usermodel"
	"github.com/gin_learn/util"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		id, err := util.ParseJwt(cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Kamu harus login dulu",
			})
			return
		}
		uId, _ := strconv.Atoi(id)
		user := usermodel.User{
			Id: uint(uId),
		}
		util.DB.Find(&user)
		c.Set("current_user", user)
		c.Next()
	}

}
