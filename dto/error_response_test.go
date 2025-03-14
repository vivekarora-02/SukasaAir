package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse_MarshalJSON(t *testing.T) {
	errResp := ErrorResponse{Error: "Test error message"}

	jsonData, err := json.Marshal(errResp)
	assert.NoError(t, err)

	expectedJSON := `{"error":"Test error message"}`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

func TestErrorResponse_UnmarshalJSON(t *testing.T) {
	jsonInput := `{"error":"Sample error"}`

	var errResp ErrorResponse
	err := json.Unmarshal([]byte(jsonInput), &errResp)
	assert.NoError(t, err)

	assert.Equal(t, "Sample error", errResp.Error)
}
