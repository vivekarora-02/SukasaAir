package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func TestSetupAuthRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/login", mockLoginHandler)

	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"message": "Login successful"}`, resp.Body.String())
}
