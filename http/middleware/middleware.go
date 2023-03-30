package middleware

import (
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Authorization string
}

func (middleware Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.Authorization = c.Request.Header.Get("Authorization")
		jwtClaim, err := common.ValidateToken(middleware.Authorization)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}

		c.Set("user", jwtClaim)
		c.Next()
	}
}
