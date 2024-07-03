package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("isLogin")
		if err != nil || cookie != "true" {
			fmt.Println("sss")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
