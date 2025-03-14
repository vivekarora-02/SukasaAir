package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockAuthMiddleware(c *gin.Context) {
	c.Next()
}

func mockReserveSeatHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Seat reserved"})
}

func mockResetSeatsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Seats reset"})
}

func TestSetupSeatRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	seatRoutes := r.Group("/seat")
	seatRoutes.Use(mockAuthMiddleware) // Bypass auth
	{
		seatRoutes.POST("/reserve", mockReserveSeatHandler)
		seatRoutes.POST("/reset", mockResetSeatsHandler)
	}

	reqReserve, _ := http.NewRequest("POST", "/seat/reserve", nil)
	wReserve := httptest.NewRecorder()
	r.ServeHTTP(wReserve, reqReserve)
	assert.Equal(t, http.StatusOK, wReserve.Code)
	assert.JSONEq(t, `{"message": "Seat reserved"}`, wReserve.Body.String())

	reqReset, _ := http.NewRequest("POST", "/seat/reset", nil)
	wReset := httptest.NewRecorder()
	r.ServeHTTP(wReset, reqReset)
	assert.Equal(t, http.StatusOK, wReset.Code)
	assert.JSONEq(t, `{"message": "Seats reset"}`, wReset.Body.String())
}
