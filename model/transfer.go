package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type (
	// TransferDestination schema for transfer destination
	TransferDestination struct {
		Type            string          `json:"type"`
		BankDetails     BankDetails     `json:"bank_details"`
		PersonalDetails PersonalDetails `json:"personal_details"`
	}

	// InitiateTransferRequest schema for initiate transfer request
	InitiateTransferRequest struct {
		CustomerID  string              `json:"customer_id"`
		Amount      float64             `json:"amount"`
		Currency    string              `json:"currency"`
		Destination TransferDestination `json:"destination"`
		Note        string              `json:"note,omitempty"`
		Reason      string              `json:"reason"`
		Reference   string              `json:"reference"`
	}

	// InitiateTerminalTransferRequest schema for initiate terminal transfer request
	InitiateTerminalTransferRequest struct {
		Amount              float64              `json:"amount"`
		SourceCurrency      string               `json:"source_currency"`
		DestinationCurrency string               `json:"destination_currency"`
		UseBalance          string               `json:"use_balance"`
		BeneficiaryID       *string              `json:"beneficiary_id,omitempty"`
		Destination         *TransferDestination `json:"destination,omitempty"`
		Note                *string              `json:"note,omitempty"`
		Reason              string               `json:"reason"`
	}

	// TransferResponse schema for transfer response
	TransferResponse struct {
		ID uuid.UUID `json:"id"`
		InitiateTransferRequest
		CreatedAt time.Time `json:"created_at"`
		Status    string    `json:"status"`
	}

	// ExchangeRateDetails schema for exchange rate details
	ExchangeRateDetails struct {
		ExchangeRate     float64 `json:"exchange_rate"`
		FeeFlat          float64 `json:"flat_fee"`
		FeePercentage    float64 `json:"fee_percentage"`
		FeeAmount        float64 `json:"fee_amount"`
		AmountReceivable float64 `json:"amount_receivable"`
	}

	// Transfer schema for customer transfer
	Transfer struct {
		ID              uuid.UUID      `json:"id"`
		Name            string         `json:"name"`
		Email           string         `json:"email"`
		CustomerID      uuid.UUID      `json:"customer_id"`
		Amount          float64        `json:"amount"`
		Currency        string         `json:"currency"`
		Destination     postgres.Jsonb `json:"destination"`
		Note            *string        `json:"note"`
		Reason          string         `json:"reason"`
		CreatedAt       time.Time      `json:"created_at"`
		CompletedAt     sql.NullTime   `json:"completed_at"`
		UpdatedAt       sql.NullTime   `json:"updated_at"`
		BatchDate       time.Time      `json:"batch_date"`
		Status          string         `json:"status"`
		Reference       string         `json:"reference"`
		CancelReason    *string        `json:"cancel_reason"`
		TransactionType string         `json:"type"`
	}

	// TerminalTransfer schema for terminal transfer
	TerminalTransfer struct {
		ID                 uuid.UUID      `json:"id"`
		BusinessID         uuid.UUID      `json:"business_id"`
		Type               string         `json:"type"`
		Amount             Money          `json:"amount"`
		Deposit            Money          `json:"deposited_amount"`
		Transfer           Money          `json:"transferred_amount"`
		SourceCurrency     string         `json:"source_currency"`
		Fee                Money          `json:"fee"`
		FeePercentage      float64        `json:"fee_percentage"`
		FeeFlat            float64        `json:"fee_flat"`
		Status             string         `json:"status"`
		ComplianceStatus   string         `json:"compliance_status"`
		BeneficiaryDetails postgres.Jsonb `json:"beneficiary_details"`
		Note               *string        `json:"note"`
		Reason             string         `json:"reason"`
		Reference          *string        `json:"reference"`
		Modified           bool           `json:"modified"`
		NeedDocumentUpload bool           `json:"need_document_upload"`
		MarkupValue        float64        `json:"markup_value"`
		CancelReason       *string        `json:"cancel_reason"`
		CompletedAt        *time.Time     `json:"completed_at"`
		CreatedAt          time.Time      `json:"created_at"`
		UpdatedAt          *time.Time     `json:"updated_at"`
		IsDateUpdated      bool           `json:"is_date_updated"`
		ComplianceNotes    *string        `json:"compliance_notes"`
	}

	// AllTransfersResponse schema for all transfers response
	AllTransfersResponse struct {
		Items []TerminalTransfer `json:"items"`
		Page  PageInfo           `json:"page"`
	}

	// Settlement schema for settlement
	Settlement struct {
		ID                uuid.UUID       `json:"id"`
		Status            string          `json:"status"`
		BatchDate         *postgres.Jsonb `json:"batch_date"`
		TransactionAmount float64         `json:"transaction_amount"`
		BatchAmount       float64         `json:"batch_amount"`
		Currency          string          `json:"currency"`
		InitiatedAt       time.Time       `json:"initiated_at"`
		CompletedTime     *time.Time      `json:"completed_at"`
		TransactionType   string          `json:"transaction_type"`
	}
)
