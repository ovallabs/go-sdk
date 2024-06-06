package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// FeeTypePercentage represent percentage fee type
	FeeTypePercentage FeeType = "percentage"

	// FeeTypeAmount represent amount fee type
	FeeTypeAmount FeeType = "amount"
)

type (
	// FeeType feeType string
	FeeType string
	// Withdrawal schema for withdrawal
	Withdrawal struct {
		ID                 uuid.UUID       `json:"id"`
		CustomerID         uuid.UUID       `json:"customer_id"`
		Name               string          `json:"name"`
		Email              string          `json:"email"`
		Reference          string          `json:"reference"`
		Amount             float64         `json:"amount"`
		Channel            string          `json:"channel"`
		Currency           string          `json:"currency"`
		CreatedAt          time.Time       `json:"created_at"`
		CompletedAt        *time.Time      `json:"completed_at"`
		UpdatedAt          *time.Time      `json:"updated_at"`
		BatchDate          string          `json:"batch_date"`
		Status             string          `json:"status"`
		WithdrawalAmount   *float64        `json:"withdrawal_amount"`
		WithdrawalCurrency *string         `json:"withdrawal_currency"`
		WithdrawalDetail   *postgres.Jsonb `json:"payout_detail"`
		CancelReason       *string         `json:"cancel_reason"`
		YieldOfferingID    uuid.UUID       `json:"yield_offering_id"`
	}

	// WithdrawalRequest schema for withdrawal request
	WithdrawalRequest struct {
		CustomerID      string  `json:"customer_id"`
		Reference       string  `json:"reference"`
		Amount          float64 `json:"amount"`
		YieldOfferingID string  `json:"yield_offering_id"`
		PayoutCurrency  *string `json:"payout_currency,omitempty"`
		WalletDetail    struct {
			Asset   string `json:"asset"`
			Network string `json:"network"`
			Address string `json:"address"`
		} `json:"wallet_detail,omitempty"`
		BankDetail struct {
			BankCode      string `json:"bank_code"`
			AccountNumber string `json:"account_number"`
		} `json:"bank_detail,omitempty"`
	}

	// FeeWithdrawalRequest schema for fee withdrawal request
	FeeWithdrawalRequest struct {
		CustomerID          string  `json:"customer_id"`
		Reference           string  `json:"reference"`
		WithdrawalReference string  `json:"withdrawal_reference"`
		Reason              string  `json:"reason"`
		FeeType             FeeType `json:"fee_type"`
		Amount              float64 `json:"amount,omitempty"`
		Percentage          float64 `json:"percentage,omitempty"`
		YieldOfferingID     string  `json:"yield_offering_id"`
	}

	// FeeWithdrawalResponse schema for fee withdrawal response
	FeeWithdrawalResponse struct {
		ID                  uuid.UUID `json:"id"`
		CustomerID          uuid.UUID `json:"customer_id"`
		Reference           string    `json:"reference"`
		WithdrawalReference string    `json:"withdrawal_reference"`
		Reason              string    `json:"reason"`
		FeeType             FeeType   `json:"fee_type"`
		Amount              float64   `json:"amount"`
		Percentage          float64   `json:"percentage"`
		YieldOfferingID     uuid.UUID `json:"yield_offering_id"`
	}
)
