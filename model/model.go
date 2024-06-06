// Package model defines object and payload models
package model

type (
	// Key is a middleware key sting value
	Key string
)

const (
	// BaseURL is the definition of ovalfi base url
	BaseURL = "https://sandbox-api.ovalfi-app.com/api/"

	// PublicKey sample sandbox environment signature
	PublicKey = "YbAO71rFXyWp0WJq-_yH7AFV6cZ7P71V53Y=" //"_Wjz3hGNJ8h1FwjJhNHnHXJJmT9Dkg=="  // "XC-WlyMxbC7MdS-mlzZ0G1tBBUXu"

	// BearerToken sample sandbox environment bearer token
	BearerToken = "eyJidXNpbmVzc0lEIjoiYjIxYTQ0YjAtYzI1Yi00NzRiLWE5ODYtOGFmNjI3MTA5YzE5IiwidXNlcklEIjoiOWVhYmJkYzQtOTg3Ny00ZDI4LTgyNTQtMTg4NjBjYWNjMDQ1Iiwia2V5IjoiUGVudGFtb25leSJ9"

	//PublicKey   = "6UCepOuO2ULaL7upafQMe3NPIeX0uNjyXZEKAw=="
	//BearerToken = "eyJidXNpbmVzc0lEIjoiM2VmMjE0NmMtMmE0Mi00ODM0LWFhMWYtMDhiMzQ1N2IwZjdlIiwidXNlcklEIjoiNWY3ZTVjY2MtY2U5MC00MDQ0LTk2NjUtYTExZjIyNjVlMWFlIiwia2V5IjoiYWJjZGVmMTIzNDU2In0="

	// LogStrRequest log string key
	LogStrRequest = "request"

	// LogStrResponse log string key
	LogStrResponse = "response"

	// LogStrParams log string key
	LogStrParams = "parameters"

	// LogStrForm
	LogStrForm = "form"

	// LogErrorCode log error_code
	LogErrorCode = "error_code"

	// RequestIDContextKey contact that holds the RequestID context key for
	RequestIDContextKey Key = "api_RequestIDContextKey"
	// RequestIDHeaderKey a constant for the request id header key
	RequestIDHeaderKey string = "X-REQUEST-ID"
)

type (
	// Money schema for money
	Money struct {
		// Currency is string value of the currency
		Currency string `json:"currency"`
		// Symbol is string value of the currency
		Symbol string ` json:"symbol" gorm:"-"`
		// Amount is the value of the amount
		Amount float64 `json:"amount"`
	}

	// Destination for a transfer
	Destination struct {
		BankDetails      BankDetails      `json:"bankDetails"`
		PersonalDetails  PersonalDetails  `json:"personalDetails"`
		IntermediaryBank IntermediaryBank `json:"intermediaryBank"`
	}

	// BankDetails schema for bank details
	BankDetails struct {
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
		RoutingNumber string `json:"routing_number,omitempty"`
		SwiftCode     string `json:"swift_code,omitempty"`
		BankName      string `json:"bank_name"`
		BankCode      string `json:"bank_code,omitempty"`
		BankBranch    string `json:"bank_branch,omitempty"`
		Country       string `json:"country"`
		City          string `json:"city,omitempty"`
		BankAddress   string `json:"bank_address,omitempty"`
		District      string `json:"district,omitempty"`
		PostalCode    string `json:"postal_code,omitempty"`
		IsWithinUS    string `json:"is_within_us"`
	}

	// PersonalDetails schema for personal details
	PersonalDetails struct {
		Name        string `json:"name"`
		Country     string `json:"country"`
		City        string `json:"city"`
		Address     string `json:"address"`
		District    string `json:"district,omitempty"`
		PostalCode  string `json:"postal_code,omitempty"`
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
	}

	// IntermediaryBank schema for intermediary bank
	IntermediaryBank struct {
		BankName    string `json:"bank_name,omitempty"`
		BankAddress string `json:"bank_address,omitempty"`
		Reference   string `json:"reference,omitempty"`
		SwiftCode   string `json:"swift_code"`
	}

	// PageInfo schema for page info
	PageInfo struct {
		Page            int64 `json:"page"`
		Size            int64 `json:"size"`
		HasNextPage     bool  `json:"has_next_age"`
		HasPreviousPage bool  `json:"has_previous_age"`
		TotalCount      int64 `json:"total_count"`
	}

	// TransferInstruction schema for transfer instruction
	TransferInstruction struct {
		IBAN               string `json:"iban"`
		SortCode           string `json:"sort_code"`
		Notes              string `json:"notes"`
		BankPhone          string `json:"bank_phone"`
		Reference          string `json:"reference"`
		SwiftCode          string `json:"swift_code"`
		BankAddress        string `json:"bank_address"`
		AccountName        string `json:"account_name"`
		AccountNumber      string `json:"account_number"`
		BankName           string `json:"bank_name"`
		BankCode           string `json:"bank_code"`
		RoutingNumber      string `json:"routing_number"`
		BeneficiaryAddress string `json:"beneficiary_address"`
	}

	// AccountDetails  schema for account details
	AccountDetails struct {
		City          string `json:"city"`
		Country       string `json:"country"`
		BankCode      string `json:"bank_code"`
		BankName      string `json:"bank_name"`
		District      string `json:"district"`
		SwiftCode     string `json:"swift_code"`
		BankBranch    string `json:"bank_branch"`
		IsWithinUS    string `json:"is_within_us"`
		PostalCode    string `json:"postal_code"`
		AccountName   string `json:"account_name"`
		BankAddress   string `json:"bank_address"`
		AccountNumber string `json:"account_number"`
		RoutingNumber string `json:"routing_number"`
	}

	// GenericResponse response wrapper
	GenericResponse struct {
		Code    int         `json:"status"`
		Data    interface{} `json:"data"`
		Message *string     `json:"message"`
		Error   *ErrorData  `json:"error"`
	}

	// WalletDetails schema for wallet details
	WalletDetails struct {
		WalletTag     *string `json:"wallet_tag,omitempty"`
		AssetType     string  `json:"asset_type,omitempty"`
		WalletAddress string  `json:"wallet_address"`
		Network       string  `json:"network"`
	}

	// TransferBeneficiaryDetails schema for transfer beneficiary details
	TransferBeneficiaryDetails struct {
		BankDetails         *BankDetails      `json:"bank_details,omitempty"`
		IntermediaryBank    *IntermediaryBank `json:"intermediary_bank,omitempty"`
		PersonalDetails     *PersonalDetails  `json:"personal_details,omitempty"`
		WalletDetails       *WalletDetails    `json:"wallet_details,omitempty"`
		FundsTransferMethod map[string]string `json:"funds_transfer_method"`
	}

	// Page schema for pagination request
	Page struct {
		Number            *int   `json:"number"`
		Size              *int   `json:"size"`
		SortBy            string `json:"sort_by"`
		SortDirectionDesc *bool  `json:"sort_direction_desc"`
	}

	// DateBetween schema for date range
	DateBetween struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
)
