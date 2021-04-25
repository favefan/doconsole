package jwt

import (
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"net/http"
	"time"

	"gitee.com/favefan/doconsole/pkg/util"
	"github.com/gin-gonic/gin"
)

// JWT is a token auth func.
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		var code int
		var data interface{}

		code = http.StatusOK
		tokenString := c.GetHeader("Access-Token") //Query("token")
		if tokenString == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := util.ParseToken(tokenString)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenFail
			}
		}

		if code != http.StatusOK {
			appG.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
