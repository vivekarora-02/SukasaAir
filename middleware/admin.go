package middleware

import (
	"net/http"
	"sukasaair/constants"
	"sukasaair/dto"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, exists := c.Get("email")
		if !exists || email != constants.ADMIN_EMAIL {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
