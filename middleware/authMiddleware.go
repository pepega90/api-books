package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin_learn/util"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		if _, err := util.ParseJwt(cookie); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Kamu harus login dulu",
			})
			return
		}
		c.Next()
	}

}
