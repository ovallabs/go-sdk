// Package model defines object and payload models
package model

const (
	// BaseURL is the definition of ovalfi base url
	BaseURL = "https://sandbox-api.ovalfi-app.com"
	// Signature sample sandbox environment signature
	Signature = "segsalerty@gmail.com"
	// BearerToken sample sandbox environment bearer token
	BearerToken = "segun"

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
		MobileNumber    string `json:"mobileNumber"`
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
)

// Data transfer objects
type (
	UpdatedAt struct {
		Time  string `json:"time"`
		Valid bool   `json:"valid"`
	}

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

	Customer struct {
		ID              string     `json:"id"`
		Name            string     `json:"customer_name"`
		MobileNumber    string     `json:"mobile_number"`
		Email           string     `json:"email"`
		Channel         string     `json:"channel"`
		Reference       string     `json:"reference"`
		YieldOfferingID string     `json:"api_yield_offering_id"`
		UpdatedAt       *UpdatedAt `json:"updated_at"`
		CreatedAt       string     `json:"created_at"`
	}

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

// Meta data for pagination
type Meta struct {
	PageCount     int `json:"page-count,omitempty"`
	ResourceCount int `json:"resource-count,omitempty"`
}
