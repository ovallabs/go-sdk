// Package example holds sample models for requests
package example

import "github.com/ovalfi/go-sdk/model"

var (
	// NewCreateCustomerRequest creates a CreateCustomerRequest struct to use as an example
	NewCreateCustomerRequest = model.CreateCustomerRequest{
		Name:            "Nonso",
		Email:           "chinonso@ovalfinance.com",
		Reference:       "ref123",
		MobileNumber:    "09080340608",
		YieldOfferingID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
	}
)
