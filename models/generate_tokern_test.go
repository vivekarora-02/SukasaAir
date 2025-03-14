package models

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"

	tokenString, err := GenerateToken(email)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, token)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, email, claims["email"])

	exp, ok := claims["exp"].(float64)
	assert.True(t, ok)
	assert.Greater(t, int64(exp), time.Now().Unix())
}
