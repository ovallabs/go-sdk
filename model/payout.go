package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	// SinglePayout represent single payout type
	SinglePayout PayoutType = "single"

	// MultiplePayout represent multiple payout type
	MultiplePayout PayoutType = "multiple"
)

type (
	// PayoutType payoutType string
	PayoutType string

	// PayoutAccount  object for payout response
	PayoutAccount struct {
		ID           uuid.UUID      `json:"id"`
		BusinessID   uuid.UUID      `json:"business_id"`
		BulkPayoutID uuid.UUID      `json:"bulk_payout_id"`
		Name         string         `json:"name"`
		Details      AccountDetails `json:"details"`
		Amount       Money          `json:"amount"`
		Status       string         `json:"status"`
		LookupInfo   string         `json:"lookup_info"`
		Remarks      string         `json:"remarks"`
		CompletedAt  *string        `json:"completed_at"`
		CreatedAt    time.Time      `json:"created_at"`
		UpdatedAt    time.Time      `json:"updated_at"`
	}

	// BulkPayoutConfig schema for payout config
	BulkPayoutConfig struct {
		Provider                 string  `json:"provider"`
		MinAmountPerPayout       float64 `json:"min_amount_per_payout"`
		MinCountOfPayout         int     `json:"min_count_of_payout"`
		MaxAmountPerPayout       float64 `json:"max_amount_per_payout"`
		MaxCountOfPayout         int     `json:"max_count_of_payout"`
		DoNameLookup             bool    `json:"do_name_lookup"`
		NamePercentageMatch      int     `json:"name_percentage_match"`
		FeePercentage            float64 `json:"fee_percentage"`
		FeeFlat                  float64 `json:"fee_flat"`
		FeeCap                   float64 `json:"fee_cap"`
		MaxPayoutPerDayPerPerson int64   `json:"max_payout_per_day_per_person"`
		AllowRecurring           bool    `json:"allow_recurring"`
	}

	// PayoutDetails for payout response
	PayoutDetails struct {
		ID           uuid.UUID  `json:"id"`
		BusinessID   uuid.UUID  `json:"business_id"`
		Status       string     `json:"status"`
		Count        int        `json:"count"`
		Currency     string     `json:"currency"`
		TotalAmount  int        `json:"total_amount"`
		Fee          Money      `json:"fee"`
		Remarks      string     `json:"remarks"`
		CancelReason *string    `json:"cancel_reason"`
		CompletedAt  *time.Time `json:"completed_at"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
	}

	// PayoutResponse get payout by ID response
	PayoutResponse struct {
		Items      PayoutDetails   `json:"items"`
		Attributes []PayoutAccount `json:"attributes"`
	}

	// AllPayoutsResponse schema for all payouts response
	AllPayoutsResponse struct {
		Items []PayoutDetails `json:"items"`
		Page  PageInfo        `json:"page"`
	}

	// CancelPayoutRequest request schema for cancel payout
	CancelPayoutRequest struct {
		BulkPayoutID string `json:"payout_id"`
		Reason       string `json:"reason"`
	}

	// InitiateBulkPayoutRequest schema
	InitiateBulkPayoutRequest struct {
		Currency        string                       `json:"currency"`
		Remarks         string                       `json:"remarks,omitempty"`
		Accounts        []BulkPayoutRecipientAccount `json:"accounts,omitempty"`
		BeneficiaryType PayoutType                   `json:"beneficiary_type"`
		BeneficiaryID   *string                      `json:"beneficiary_id,omitempty"`
		Amount          *float64                     `json:"amount,omitempty"`
	}

	// BulkPayoutRecipientAccount schema
	BulkPayoutRecipientAccount struct {
		Amount      float64                    `json:"amount"`
		Destination TransferBeneficiaryDetails `json:"destination"`
		Remarks     string                     `json:"remarks"`
	}
)
