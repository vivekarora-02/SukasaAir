package routes

import (
	"sukasaair/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine) *gin.Engine {
	r.POST("/login", controllers.LoginHandler)

	return r
}
