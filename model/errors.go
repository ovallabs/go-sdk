package model

import "errors"

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")
)
