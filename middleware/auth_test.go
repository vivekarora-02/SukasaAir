package middleware

import (
	"net/http"
	"net/http/httptest"
	"sukasaair/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(method, path, token string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(AuthMiddleware)
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	req := httptest.NewRequest(method, path, nil)
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	token, _ := models.GenerateToken("user@example.com")
	w := performRequest("GET", "/protected", token)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Success")
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	w := performRequest("GET", "/protected", "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Missing token")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	w := performRequest("GET", "/protected", "invalid.token.string")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token")
}
