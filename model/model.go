// Package model defines object and payload models
package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	// BaseURL is the definition of ovalfi base url
	BaseURL = "https://sandbox-api.ovalfi-app.com/api/"

	// PublicKey sample sandbox environment signature
	PublicKey = "_Wjz3hGNJ8h1FwjJhNHnHXJJmT9Dkg=="

	// BearerToken sample sandbox environment bearer token
	BearerToken = "eyJidXNpbmVzc0lEIjoiOTIzYjJkZjUtNGE4OS00Y2ViLWIxNDgtYzJlNWFjNTJkMDRlIiwidXNlcklEIjoiMjQ4YmFhNDMtYzQ0Yi00ZjYwLWI2MWQtY2VlZjYwOThjNzg1Iiwia2V5IjoidXBwcHBwIn0="

	// LogStrRequest log string key
	LogStrRequest = "request"

	// LogStrResponse log string key
	LogStrResponse = "response"

	// LogErrorCode log error_code
	LogErrorCode = "error_code"

	// Credit action credit a user from it holding balance to a yield offering
	Credit FundTransferAction = "credit"

	// Debit action debit a user from it holding balance to a yield offering
	Debit FundTransferAction = "debit"
)

// Requests for the endpoints
type (
	// CreateCustomerRequest attributes payload to create new API customer
	CreateCustomerRequest struct {
		Name             string      `json:"name"`
		Email            string      `json:"email"`
		Reference        string      `json:"reference"`
		MobileNumber     string      `json:"mobile_number"`
		YieldOfferingIDs []uuid.UUID `json:"yield_offering_ids"`
	}

	// UpdateCustomerRequest attributes payload to update API customer
	UpdateCustomerRequest struct {
		CustomerID       string      `json:"customer_id"`
		Name             string      `json:"name"`
		Email            string      `json:"email"`
		Reference        string      `json:"reference"`
		MobileNumber     string      `json:"mobile_number"`
		YieldOfferingIDs []uuid.UUID `json:"yield_offering_ids"`
	}

	// GetCustomerByIDRequest attributes payload to update API customer
	GetCustomerByIDRequest struct {
		CustomerID string `json:"customer_id"`
	}

	// GetCustomerBalanceRequest attributes payload to get customer balance
	GetCustomerBalanceRequest struct {
		CustomerID      uuid.UUID `json:"customer_id"`
		YieldOfferingID uuid.UUID `json:"yield_offering_id"`
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

	// UpdateYieldOfferingProfilesRequest payload for updating yield offerings
	UpdateYieldOfferingProfilesRequest struct {
		YieldOfferingID string `json:"yield_offering_id"`
		Name            string `json:"name"`
		Description     string `json:"description"`
	}

	// GetYieldProfileByIDRequest attributes payload to update yield offering by ID
	GetYieldProfileByIDRequest struct {
		YieldProfileID string `json:"yield_offering_id"`
	}

	// InitiateDepositRequest attributes payload to initiate a new API deposit
	InitiateDepositRequest struct {
		CustomerID string  `json:"customer_id"`
		Reference  string  `json:"reference"`
		Amount     float64 `json:"amount"`
	}

	// TransactionRequest attributes payload for getting transactions
	TransactionRequest struct {
		CustomerID      *uuid.UUID `json:"customer_id"`
		BatchDate       *string    `json:"batch_date"`
		Reference       *string    `json:"reference"`
		YieldOfferingID *uuid.UUID `json:"yield_offering_id"`
		Size            *int       `json:"size"`
		Page            *int       `json:"page"`
	}

	// InitiateTransferRequest attributes payload to initiate a new API transfer
	InitiateTransferRequest struct {
		CustomerID  string              `json:"customer_id"`
		Amount      float64             `json:"amount"`
		Currency    string              `json:"currency"`
		Destination TransferDestination `json:"destination"`
		Note        string              `json:"note"`
		Reason      string              `json:"reason"`
		Reference   string              `json:"reference"`
	}

	// TransferDestination holds recipient's bank and personal info
	TransferDestination struct {
		BankDetails     BankDetails     `json:"bank_details"`
		PersonalDetails PersonalDetails `json:"personal_details"`
	}

	// InitiateWithdrawalRequest attributes payload to initiate a new API withdrawal
	InitiateWithdrawalRequest struct {
		BusinessID string  `json:"business_id"`
		CustomerID string  `json:"customer_id"`
		Reference  string  `json:"reference"`
		Amount     float64 `json:"amount"`
	}

	// WalletRequest attributes payload to get wallet address
	WalletRequest struct {
		CustomerID string `json:"customer_id"`
		Network    string `json:"network"`
		Asset      string `json:"asset"`
	}

	FundTransferAction string

	// FundTransferRequest attributes payload to transfer funds from one yield offering to another
	FundTransferRequest struct {
		CustomerID      uuid.UUID          `json:"customer_id"`
		Reference       string             `json:"reference"`
		Amount          float64            `json:"amount"`
		Action          FundTransferAction `json:"action"`
		YieldOfferingID uuid.UUID          `json:"yield_offering_id"`
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
		ID              string      `json:"id"`
		BusinessID      string      `json:"businessID"`
		CustomerID      string      `json:"customerID"`
		YieldOfferingID string      `json:"yieldOfferingID"`
		Type            string      `json:"type"`
		Amount          float64     `json:"amount"`
		Currency        string      `json:"currency"`
		Reference       string      `json:"reference"`
		Status          string      `json:"status"`
		Destination     Destination `json:"destination"`
		CompletedAt     string      `json:"completedAt"`
		CreatedAt       string      `json:"createdAt"`
		BatchDate       string      `json:"batchDate"`
	}

	// Customer data object
	Customer struct {
		ID               string       `json:"id"`
		Name             string       `json:"customer_name"`
		MobileNumber     string       `json:"mobile_number"`
		Email            string       `json:"email"`
		Channel          string       `json:"channel"`
		Reference        string       `json:"reference"`
		YieldOfferingIDs []uuid.UUID  `json:"api_yield_offering_ids"`
		UpdatedAt        sql.NullTime `json:"updated_at"`
		CreatedAt        string       `json:"created_at"`
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

	// Portfolio data object for business portfolios
	Portfolio struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Network     string `json:"network"`
		APYRate     string `json:"apy_rate"`
	}

	// YieldOfferingProfile data object for yield offerings
	YieldOfferingProfile struct {
		YieldOfferingID       string  `json:"yield_offering_id"`
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

	// UpdatedYieldOfferingProfile data object for updated yield offerings
	UpdatedYieldOfferingProfile struct {
		YieldOfferingID string       `json:"yield_offering_id"`
		Name            string       `json:"name"`
		Description     string       `json:"description"`
		CreatedAt       string       `json:"created_at"`
		UpdatedAt       sql.NullTime `json:"updated_at"`
		Reference       string       `json:"reference"`
	}

	// DepositBatchResponse to get deposit batch
	DepositBatchResponse struct {
		Deposits           map[string]DepositResponse   `json:"deposits"`
		TotalAmount        float64                      `json:"total_amount"`
		PaidDepositDetails []*ExternalAPIDepositDetails `json:"paid_deposit_details"`
	}
	// DepositResponse as response payload for settled/unsettled deposit payment
	DepositResponse struct {
		Deposits    []*Deposit `json:"deposits"`
		TotalAmount float64    `json:"total_amount"`
	}
	// ExternalAPIDepositDetails struct gives details about the external api deposit
	ExternalAPIDepositDetails struct {
		From             string  `json:"from"`
		To               string  `json:"to"`
		TotalAmount      float64 `json:"total_amount"`
		AmountPaid       float64 `json:"amount_paid"`
		BalanceRemaining float64 `json:"balance_remaining"`
	}

	// Deposit data objet
	Deposit struct {
		ID              uuid.UUID  `json:"id"`
		CustomerID      uuid.UUID  `json:"customer_id"`
		BusinessID      uuid.UUID  `json:"business_id"`
		Name            string     `json:"name"`
		Email           string     `json:"email"`
		Reference       string     `json:"reference"`
		Currency        string     `json:"currency"`
		Amount          float64    `json:"amount"`
		Channel         string     `json:"channel"`
		CreatedAt       time.Time  `json:"created_at"`
		SettledAt       *time.Time `json:"settled_at"`
		BalanceBefore   float64    `json:"balance_before"`
		BalanceAfter    float64    `json:"balance_after"`
		DepositBeforeID uuid.UUID  `json:"deposit_before_id"`
		BatchDate       *time.Time `json:"batch_date"`
		Status          string     `json:"status"`
		CancelReason    *string    `json:"cancel_reason"`
	}

	// Transfer data object
	Transfer struct {
		ID uuid.UUID `json:"id"`
		InitiateTransferRequest
		CreatedAt time.Time `json:"created_at"`
		Status    string    `json:"status"`
	}

	// Withdrawal data object
	Withdrawal struct {
		ID         uuid.UUID `json:"id"`
		BusinessID uuid.UUID `json:"business_id"`
		CustomerID uuid.UUID `json:"customer_id"`
		Reference  string    `json:"reference"`
		Amount     float64   `json:"amount"`
		Status     string    `json:"status"`
		CreatedAt  time.Time `json:"created_at"`
		Channel    string    `json:"channel"`
	}

	// Wallet data object
	Wallet struct {
		CustomerID    uuid.UUID  `json:"customer_id"`
		WalletAddress string     `json:"wallet_address"`
		Asset         string     `json:"asset"`
		Network       string     `json:"network"`
		Type          string     `json:"type"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
	}

	// SupportedAsset data object
	SupportedAsset struct {
		Asset    string   `json:"asset"`
		Networks []string `json:"networks"`
	}

	// CustomerBalanceResponse object
	CustomerBalanceResponse struct {
		YieldOfferingID uuid.UUID `json:"yield_offering_id"`
		Name            string    `json:"name"`
		Currency        string    `json:"currency"`
		Amount          float64   `json:"balance"`
	}

	// CustomerBalancesResponse object
	CustomerBalancesResponse struct {
		CustomerID   uuid.UUID                  `json:"customer_id"`
		TotalBalance float64                    `json:"total_balance"`
		Detail       []*CustomerBalanceResponse `json:"detail"`
	}

	// Page object
	Page struct {
		Page            int64 `json:"page"`
		Size            int64 `json:"size"`
		HasNextPage     bool  `json:"hasNextPage"`
		HasPreviousPage bool  `json:"hasPreviousPage"`
		TotalCount      int64 `json:"totalCount"`
	}

	// TransactionResponse object
	TransactionResponse struct {
		Item struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"item"`
		Page Page `json:"page"`
	}
)
