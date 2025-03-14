package controllers

import (
	"net/http"
	"sukasaair/dto"
	"sukasaair/models"

	"github.com/gin-gonic/gin"
)

// LoginHandler handles user authentication
// @Summary User Login
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "User login request body"
// @Success 200 {object} dto.LoginResponse "Successful login"
// @Failure 400 {object} dto.ErrorResponse "Invalid request format"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request format"})
		return
	}

	token, err := models.GenerateToken(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
