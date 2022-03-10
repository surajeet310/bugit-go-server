package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajeet310/bugit-go-server/users"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := users.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
