package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Credit represent credit transfer action
	Credit FundTransferAction = "credit"

	// Debit represent debit transfer action
	Debit FundTransferAction = "debit"
)

type (
	// Deposit schema for deposit
	Deposit struct {
		ID                uuid.UUID  `json:"id"`
		CustomerID        uuid.UUID  `json:"customer_id"`
		BusinessID        uuid.UUID  `json:"business_id"`
		Name              string     `json:"name"`
		Email             string     `json:"email"`
		Reference         string     `json:"reference"`
		Currency          string     `json:"currency"`
		Amount            float64    `json:"amount"`
		AmountDeposited   float64    `json:"deposited_amount"`
		DepositedCurrency string     `json:"deposited_currency"`
		Channel           string     `json:"channel"`
		CreatedAt         time.Time  `json:"created_at"`
		SettledAt         *time.Time `json:"settled_at"`
		BalanceBefore     float64    `json:"balance_before"`
		BalanceAfter      float64    `json:"balance_after"`
		DepositBeforeID   uuid.UUID  `json:"deposit_before_id"`
		Status            string     `json:"status"`
		CancelReason      *string    `json:"cancel_reason"`
		YieldOfferingID   uuid.UUID  `json:"yield_offering_id"`
	}

	// InitiateDepositRequest schema for initiate deposit request
	InitiateDepositRequest struct {
		CustomerID      string  `json:"customer_id"`
		Reference       string  `json:"reference"`
		Amount          float64 `json:"amount"`
		YieldOfferingID string  `json:"yield_offering_id"`
	}

	// DepositBatchResponse schema for deposit batch response
	DepositBatchResponse struct {
		Deposits map[string]struct {
			Deposits    []*Deposit `json:"deposits"`
			TotalAmount float64    `json:"total_amount"`
		} `json:"deposits"`
		TotalAmount float64 `json:"total_amount"`
	}

	// FundTransferAction transferAction type string
	FundTransferAction string

	// FundTransferRequest schema for func transfer request
	FundTransferRequest struct {
		CustomerID      string             `json:"customer_id"`
		Reference       string             `json:"reference"`
		Amount          float64            `json:"amount"`
		Action          FundTransferAction `json:"action"`
		YieldOfferingID string             `json:"yield_offering_id"`
	}

	// TransferParty schema for transfer response
	TransferParty struct {
		CustomerID      string `json:"customer_id"`
		YieldOfferingID string `json:"yield_offering_id"`
	}

	// IntraTransferRequest schema for intra transfer request
	IntraTransferRequest struct {
		Reference string        `json:"reference"`
		Amount    float64       `json:"amount"`
		Sender    TransferParty `json:"sender"`
		Receiver  TransferParty `json:"receiver"`
	}

	// IntraTransferResponse schema for intra transfer response
	IntraTransferResponse struct {
		ID        uuid.UUID     `json:"id"`
		Reference string        `json:"reference"`
		Amount    float64       `json:"amount"`
		Sender    TransferParty `json:"sender"`
		Receiver  TransferParty `json:"receiver"`
	}
)
