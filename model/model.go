// Package model defines object and payload models
package model

const (
	// BaseURL is the definition of ovalfi base url
	BaseURL = "https://sandbox-api.ovalfi-app.com"
	// Signature sample sandbox environment signature
	Signature = "segsalerty@gmail.com"
	// BearerToken sample sandbox environment bearer token
	BearerToken = "segun"
)

type (
	ResponsePayload struct {
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)
