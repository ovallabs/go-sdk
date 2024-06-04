// Package model defines object and payload models
package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

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

	// Credit action credit a user from it holding balance to a yield offering
	Credit FundTransferAction = "credit"

	// Debit action debit a user from it holding balance to a yield offering
	Debit FundTransferAction = "debit"

	// FeeTypePercentage represent FeeType in percentage
	FeeTypePercentage FeeType = "percentage"

	// FeeTypeAmount represent FeeType in amount
	FeeTypeAmount FeeType = "amount"

	// RequestIDContextKey contact that holds the RequestID context key for
	RequestIDContextKey Key = "api_RequestIDContextKey"
	// RequestIDHeaderKey a constant for the request id header key
	RequestIDHeaderKey string = "X-REQUEST-ID"
)

type (
	// Money struct
	Money struct {
		// Currency is string value of the currency
		Currency string `json:"currency"`
		// Symbol is string value of the currency
		Symbol string ` json:"symbol" gorm:"-"`
		// Amount is the value of the amount
		Amount float64 `json:"amount"`
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

	// FundTransferAction string
	FundTransferAction string

	// FundTransferRequest attributes payload to transfer funds from one yield offering to another
	FundTransferRequest struct {
		CustomerID      uuid.UUID          `json:"customer_id"`
		Reference       string             `json:"reference"`
		Amount          float64            `json:"amount"`
		Action          FundTransferAction `json:"action"`
		YieldOfferingID uuid.UUID          `json:"yield_offering_id"`
	}

	// TransferParty attributes payload to hold sender and receiver payload
	TransferParty struct {
		CustomerID      string `json:"customer_id"`
		YieldOfferingID string `json:"yield_offering_id"`
	}

	// IntraTransferRequest attributes payload to transfer funds between customers
	IntraTransferRequest struct {
		Reference string        `json:"reference"`
		Amount    float64       `json:"amount"`
		Sender    TransferParty `json:"sender"`
		Receiver  TransferParty `json:"receiver"`
	}

	// BankDetail bank details for withdrawal
	BankDetail struct {
		BankCode      string `json:"bank_code"`
		AccountNumber string `json:"account_number"`
	}

	// WalletDetail wallet details for withdrawal
	WalletDetail struct {
		Address string `json:"address"`
		Network string `json:"network"`
		Asset   string `json:"asset"`
	}

	// WithdrawalRequest attribute payload for crypto and fiat withdrawal
	WithdrawalRequest struct {
		CustomerID      uuid.UUID     `json:"customer_id"`
		Reference       string        `json:"reference"`
		Amount          float64       `json:"amount"`
		YieldOfferingID uuid.UUID     `json:"yield_offering_id"`
		PayoutCurrency  string        `json:"payout_currency"`
		BankDetail      *BankDetail   `json:"bank_detail"`
		WalletDetail    *WalletDetail `json:"wallet_detail"`
	}

	// FeeWithdrawalRequest attribute payload for fee withdrawal
	FeeWithdrawalRequest struct {
		ID                  uuid.UUID `json:"id"`
		CustomerID          uuid.UUID `json:"customer_id" validate:"required"`
		BusinessID          uuid.UUID `json:"-"`
		Reference           string    `json:"reference" validate:"required"`
		WithdrawalReference string    `json:"withdrawal_reference" validate:"required"`
		FeeType             FeeType   `json:"fee_type" validate:"required,oneof=amount percentage"`
		Amount              float64   `json:"amount,omitempty" validate:"required_if=FeeType amount"`
		Percentage          float64   `json:"percentage,omitempty" validate:"required_if=FeeType percentage"`
		YieldOfferingID     uuid.UUID `json:"yield_offering_id" validate:"required"`
		Reason              string    `json:"reason" validate:"required"`
	}

	// FeeType feeType string
	FeeType string
)

type (
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
		PhoneNumber string `json:"phone_number,omitempty"`
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
		DepositBeforeID uuid.UUID  `json:"deposit_before_id"`
		BatchDate       *string    `json:"batch_date"`
		Status          string     `json:"status"`
		CancelReason    *string    `json:"cancel_reason"`
	}

	// Withdrawal data object
	Withdrawal struct {
		ID                 string      `json:"id"`
		CustomerID         string      `json:"customer_id"`
		Reference          string      `json:"reference"`
		Amount             float64     `json:"amount"`
		Channel            string      `json:"channel"`
		Currency           string      `json:"currency"`
		CreatedAt          time.Time   `json:"created_at"`
		CompletedAt        *time.Time  `json:"completed_at"`
		UpdatedAt          *time.Time  `json:"updated_at"`
		BatchDate          *string     `json:"batch_date"`
		Status             string      `json:"status"`
		WithdrawalAmount   float64     `json:"withdrawal_amount"`
		WithdrawalCurrency string      `json:"withdrawal_currency"`
		PayoutDetail       interface{} `json:"payout_detail"`
		CancelReason       *string     `json:"cancel_reason"`
		YieldOfferingID    string      `json:"yield_offering_id"`
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
		LogoURL  string   `json:"logo_url"`
		Networks []string `json:"networks"`
	}

	// PageInfo object
	PageInfo struct {
		Page            int64 `json:"page"`
		Size            int64 `json:"size"`
		HasNextPage     bool  `json:"has_next_age"`
		HasPreviousPage bool  `json:"has_previous_age"`
		TotalCount      int64 `json:"total_count"`
	}

	// TransactionResponse object
	TransactionResponse struct {
		Items struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"items"`
		Page PageInfo `json:"page"`
	}

	// IntraTransferResponse response payload for intra transfer
	IntraTransferResponse struct {
		ID        uuid.UUID     `json:"id"`
		Reference string        `json:"reference"`
		Amount    float64       `json:"amount"`
		Sender    TransferParty `json:"sender"`
		Receiver  TransferParty `json:"receiver"`
	}

	// TransferInstruction for transfer instruction details
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

	// FeeWithdrawal object for fee withdrawal
	FeeWithdrawal struct {
		ID                  uuid.UUID `json:"id"`
		CustomerID          uuid.UUID `json:"customer_id" validate:"required"`
		BusinessID          uuid.UUID `json:"-"`
		Reference           string    `json:"reference" validate:"required"`
		WithdrawalReference string    `json:"withdrawal_reference" validate:"required"`
		FeeType             FeeType   `json:"fee_type" validate:"required,oneof=amount percentage"`
		Amount              float64   `json:"amount,omitempty" validate:"required_if=FeeType amount"`
		Percentage          float64   `json:"percentage,omitempty" validate:"required_if=FeeType percentage"`
		YieldOfferingID     uuid.UUID `json:"yield_offering_id" validate:"required"`
		Reason              string    `json:"reason" validate:"required"`
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

	// TransferBeneficiaryDetails  request schema for update payout
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
