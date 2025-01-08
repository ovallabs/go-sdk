package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	// GenerateBankAccountRequest schema for generate bank account request
	GenerateBankAccountRequest struct {
		CustomerID string `json:"customer_id"`
		Currency   string `json:"currency"`
		Reference  string `json:"reference"`

		BVN         *string `json:"bvn,omitempty"`
		PhoneNumber *string `json:"phone_number,omitempty"`

		DocumentType      *string `json:"document_type,omitempty"`
		Number            *string `json:"document_number,omitempty"`
		IssuedCountryCode *string `json:"issued_country_code,omitempty"`
		IssuedBy          *string `json:"issued_by,omitempty"`
		IssuedDate        *string `json:"issued_date,omitempty"`
		ExpirationDate    *string `json:"expiration_date,omitempty"`
		Country           *string `json:"country,omitempty"`
		ZipCode           *string `json:"zip_code,omitempty"`
		City              *string `json:"city,omitempty"`
		Street            *string `json:"street,omitempty"`
		State             *string `json:"state,omitempty"`
		DateOfBirth       *string `json:"date_of_birth,omitempty"`
	}

	// BankAccount schema for bank account
	BankAccount struct {
		ID                uuid.UUID           `json:"id"`
		CustomerID        uuid.UUID           `json:"customer_id"`
		BusinessID        uuid.UUID           `json:"business_id"`
		AccountID         string              `json:"account_id"`
		AccountDetails    TransferInstruction `json:"account_details"`
		BusinessReference string              `json:"reference"`
		Currency          string              `json:"currency"`
		CreatedAt         time.Time           `json:"created_at"`
		UpdatedTime       *time.Time          `json:"updated_at"`
	}

	// BankCode schema for bank code
	BankCode struct {
		BankName string `json:"name"`
		Code     string `json:"code"`
	}

	// Bank schema for bank
	Bank struct {
		Name    string `json:"name"`
		Code    string `json:"code"`
		Country string `json:"country"`
	}

	// AccountResolveRequest schema for account resolve request
	AccountResolveRequest struct {
		BankCode      string `json:"bank_code"`
		AccountNumber string `json:"account_number"`
	}

	// MockCustomerDepositRequest schema for customer mock deposit request
	MockCustomerDepositRequest struct {
		CustomerID string  `json:"customer_id"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
	}
)
