package controllers

import (
	"net/http"
	"sukasaair/dto"
	"sukasaair/models"

	"github.com/gin-gonic/gin"
)

// ReserveSeatHandler reserves a seat for a passenger
// @Summary Reserve a seat
// @Description Reserves a specific seat for a passenger
// @Tags Seats
// @Accept json
// @Produce json
// @Param reserveSeatRequest body dto.ReserveSeatRequest true "Seat reservation details"
// @Success 200 {object} dto.ReserveSeatResponse "Seat reserved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request format"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /seat/reserve [post]
func ReserveSeatHandler(c *gin.Context) {
	var req dto.ReserveSeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request"})
		return
	}

	err := models.ReserveSeat(req.SeatNumber, req.PassengerPhone, req.PassengerName, req.PassengerAge)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ReserveSeatResponse{Message: "Seat reserved successfully"})
}

// ResetSeatsHandler resets all seat reservations (Admin only)
// @Summary Reset all seat reservations
// @Description Resets all reserved seats (Admin access only)
// @Tags Seats
// @Produce json
// @Success 200 {object} dto.ResetSeatsResponse "All seat reservations reset successfully"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized access"
// @Security BearerAuth
// @Router /seat/reset [post]
func ResetSeatsHandler(c *gin.Context) {
	models.ResetSeats()
	c.JSON(http.StatusOK, dto.ResetSeatsResponse{Message: "All seat reservations reset successfully"})
}
