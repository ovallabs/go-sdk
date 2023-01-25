package model

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNetworkError when something goes wrong with the API call
	ErrNetworkError = errors.New("network error")
)

// ParseError parses error message to one that we can see publicly
func ParseError(errMsg error) error {
	// the error looks like: "user with this email does not exist"
	response := fmt.Errorf("%q%q", strings.ToUpper(errMsg.Error()[0:1]), errMsg.Error()[1:])
	return response
}
