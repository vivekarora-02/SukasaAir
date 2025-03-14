package dto

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginRequestValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		input      LoginRequest
		shouldPass bool
	}{
		{"Valid Email", LoginRequest{Email: "user@example.com"}, true},
		{"Missing Email", LoginRequest{Email: ""}, false},
		{"Invalid Email", LoginRequest{Email: "invalid-email"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.input)
			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = req
			var requestBody LoginRequest
			err := c.ShouldBindJSON(&requestBody)

			if tt.shouldPass {
				assert.NoError(t, err, "Expected valid input to pass")
			} else {
				assert.Error(t, err, "Expected invalid input to fail")
			}
		})
	}
}
