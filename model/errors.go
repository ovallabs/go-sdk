package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")
)

// ParseError parses error message to a more specific format
func ParseError(errMsg string) error {
	response := fmt.Errorf("%s", errors.New(errMsg))
	return response
}

type (
	// ErrorResponse object
	ErrorResponse struct {
		Status  int       `json:"status"`
		Data    string    `json:"data"`
		Message string    `json:"message"`
		Error   ErrorData `json:"error"`
	}
	// ErrorData struct
	ErrorData struct {
		ID      string `json:"id"`
		Details string `json:"details"`
		Message string `json:"message"`
	}
)

// GetErrorDetails to unmarshal the err response gotten from api-service
func GetErrorDetails(errMsg string) (ErrorResponse, error) {
	var result ErrorResponse
	err := json.Unmarshal([]byte(errMsg), &result)
	if err != nil {
		log.Err(err).Msg("json.Unmarshal failed")
		return ErrorResponse{}, err
	}
	return result, nil
}
