// Package model defines object and payload models
package model

import "database/sql"

const (
	// BaseURL is the definition of ovalfi base url
	BaseURL = "https://sandbox-api.ovalfi-app.com/api/"

	// PublicKey sample sandbox environment signature
	PublicKey = "IYtqe0xG0voYzbPhUEtTIEKyKj4Keq0O"

	// BearerToken sample sandbox environment bearer token
	BearerToken = "eyJidXNpbmVzc0lEIjoiOTIzYjJkZjUtNGE4OS00Y2ViLWIxNDgtYzJlNWFjNTJkMDRlIiwidXNlcklEIjoiMjQ4YmFhNDMtYzQ0Yi00ZjYwLWI2MWQtY2VlZjYwOThjNzg1Iiwia2V5Ijoib2xhcHJvZzEifQ=="

	// LogStrRequest log string key
	LogStrRequest = "request"

	// LogStrResponse log string key
	LogStrResponse = "response"
)

// Requests for the endpoints
type (
	// CreateCustomerRequest attributes payload to create new API customer
	CreateCustomerRequest struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Reference       string `json:"reference"`
		MobileNumber    string `json:"mobile_number"`
		YieldOfferingID string `json:"yield_offering_id"`
	}

	// UpdateCustomerRequest attributes payload to update API customer
	UpdateCustomerRequest struct {
		CustomerID      string `json:"customer_id"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		Reference       string `json:"reference"`
		MobileNumber    string `json:"mobile_number"`
		YieldOfferingID string `json:"yield_offering_id"`
	}

	// GetCustomerByIDRequest attributes payload to update API customer
	GetCustomerByIDRequest struct {
		CustomerID string `json:"customer_id"`
	}

	// CreateYieldOfferingProfilesRequest payload for API yield offerings
	CreateYieldOfferingProfilesRequest struct {
		Name                  string  `json:"name"`
		Description           string  `json:"description"`
		APYRate               float64 `json:"apy_rate"`
		Currency              string  `json:"currency"`
		DepositLockDay        int     `json:"deposit_lock_day"`
		MinimumDepositAllowed float64 `json:"minimum_deposit_allowed"`
		MaximumDepositAllowed float64 `json:"maximum_deposit_allowed"`
		YieldableAfterDay     int     `json:"yieldable_after_day"`
		WithdrawalLimitRate   float64 `json:"withdrawal_limit_rate"`
		PortfolioID           string  `json:"portfolio_id"`
		Reference             string  `json:"reference"`
	}
)

type (
	// Destination for a transfer
	Destination struct {
		BankDetails      BankDetails      `json:"bankDetails"`
		PersonalDetails  PersonalDetails  `json:"personalDetails"`
		IntermediaryBank IntermediaryBank `json:"intermediaryBank"`
	}

	// BankDetails recipient's bank details
	BankDetails struct {
		AccountNumber string `json:"accountNumber"`
		AccountName   string `json:"accountName"`
		RoutingNumber string `json:"routingNumber"`
		SwiftCode     string `json:"swiftCode"`
		BankName      string `json:"bankName"`
		BankBranch    string `json:"bankBranch"`
		Country       string `json:"country"`
		City          string `json:"city"`
		BankAddress   string `json:"bankAddress"`
		District      string `json:"district"`
		PostalCode    string `json:"postalCode"`
		IsWithinUS    string `json:"isWithinUS"`
	}

	// PersonalDetails recipient's personal details
	PersonalDetails struct {
		Name       string `json:"name"`
		Country    string `json:"country"`
		City       string `json:"city"`
		Address    string `json:"address"`
		District   string `json:"district"`
		PostalCode string `json:"postalCode"`
	}

	// IntermediaryBank recipient's intermediary bank
	IntermediaryBank struct {
		BankName    string `json:"bankName"`
		BankAddress string `json:"bankAddress"`
		Reference   string `json:"reference"`
		SwiftCode   string `json:"swiftCode"`
	}

	// Transaction data object for customer transactions
	Transaction struct {
		ID          string      `json:"id"`
		BusinessID  string      `json:"businessID"`
		CustomerID  string      `json:"customerID"`
		Type        string      `json:"type"`
		Amount      float64     `json:"amount"`
		Currency    string      `json:"currency"`
		Reference   string      `json:"reference"`
		Status      string      `json:"status"`
		Destination Destination `json:"destination"`
		CompletedAt string      `json:"completedAt"`
		CreatedAt   string      `json:"createdAt"`
		BatchDate   string      `json:"batchDate"`
	}

	// Customer data object
	Customer struct {
		ID              string       `json:"id"`
		Name            string       `json:"customer_name"`
		MobileNumber    string       `json:"mobile_number"`
		Email           string       `json:"email"`
		Channel         string       `json:"channel"`
		Reference       string       `json:"reference"`
		YieldOfferingID string       `json:"api_yield_offering_id"`
		UpdatedAt       sql.NullTime `json:"updated_at"`
		CreatedAt       string       `json:"created_at"`
	}

	// CustomerInfo data object for additional customer details
	CustomerInfo struct {
		Customer
		DepositCount    int64          `json:"deposit_count"`
		TotalDeposit    float64        `json:"total_deposit"`
		WithdrawalCount int64          `json:"withdrawal_count"`
		TotalWithdrawal float64        `json:"total_withdrawal"`
		Transfer        float64        `json:"transfer"`
		Balance         float64        `json:"balance"`
		Transaction     []*Transaction `json:"transaction"`
	}
)
