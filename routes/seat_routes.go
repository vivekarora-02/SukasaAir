package routes

import (
	"sukasaair/controllers"
	"sukasaair/middleware"

	"github.com/gin-gonic/gin"
)

func SetupSeatRouter(r *gin.Engine) *gin.Engine {
	seatRoutes := r.Group("/seat")
	seatRoutes.Use(middleware.AuthMiddleware)

	seatRoutes.POST("/reserve", controllers.ReserveSeatHandler)

	seatRoutes.POST("/reset", middleware.AdminMiddleware(), controllers.ResetSeatsHandler)

	return r
}
