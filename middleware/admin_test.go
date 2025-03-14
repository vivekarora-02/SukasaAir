package middleware

import (
	"net/http"
	"net/http/httptest"
	"sukasaair/constants"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdminMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		email        interface{}
		expectedCode int
	}{
		{
			name:         "Unauthorized - No email in context",
			email:        nil,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Unauthorized - Email is not admin",
			email:        "user@example.com",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Authorized - Admin email",
			email:        constants.ADMIN_EMAIL,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(func(c *gin.Context) {
				if tt.email != nil {
					c.Set("email", tt.email)
				}
				c.Next()
			})
			router.Use(AdminMiddleware())

			router.GET("/admin-test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Success"})
			})

			req, _ := http.NewRequest(http.MethodGet, "/admin-test", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusUnauthorized {
				assert.Contains(t, w.Body.String(), "Unauthorized")
			}
		})
	}
}
