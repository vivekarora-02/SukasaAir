package dto

// LoginRequest represents the request body for login
type LoginRequest struct {
	Email string `json:"emailId" binding:"required,email"`
}

// LoginResponse represents the response for login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
