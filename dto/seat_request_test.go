package dto

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReserveSeatRequestValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		input      ReserveSeatRequest
		shouldPass bool
	}{
		{"Valid Request", ReserveSeatRequest{SeatNumber: 50, PassengerPhone: "9876543210", PassengerName: "John Doe", PassengerAge: 25}, true},
		{"Missing SeatNumber", ReserveSeatRequest{PassengerPhone: "9876543210", PassengerName: "John Doe", PassengerAge: 25}, false},
		{"Invalid SeatNumber (Out of Range)", ReserveSeatRequest{SeatNumber: 500, PassengerPhone: "9876543210", PassengerName: "John Doe", PassengerAge: 25}, false},
		{"Missing PassengerPhone", ReserveSeatRequest{SeatNumber: 50, PassengerName: "John Doe", PassengerAge: 25}, false},
		{"Missing PassengerName", ReserveSeatRequest{SeatNumber: 50, PassengerPhone: "9876543210", PassengerAge: 25}, false},
		{"Missing PassengerAge", ReserveSeatRequest{SeatNumber: 50, PassengerPhone: "9876543210", PassengerName: "John Doe"}, false},
		{"Invalid PassengerAge (0 years)", ReserveSeatRequest{SeatNumber: 50, PassengerPhone: "9876543210", PassengerName: "John Doe", PassengerAge: 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.input)
			req := httptest.NewRequest("POST", "/reserve-seat", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Bind JSON to struct
			c.Request = req
			var requestBody ReserveSeatRequest
			err := c.ShouldBindJSON(&requestBody)

			// Validate response
			if tt.shouldPass {
				assert.NoError(t, err, "Expected valid input to pass")
			} else {
				assert.Error(t, err, "Expected invalid input to fail")
			}
		})
	}
}
