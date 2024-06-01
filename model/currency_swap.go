package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	// InitiateCurrencySwapRequest schema for currency swap request
	InitiateCurrencySwapRequest struct {
		FromCurrency string  `json:"from_currency"`
		ToCurrency   string  `json:"to_currency"`
		Amount       float64 `json:"amount"`
	}

	// CurrencySwap schema for currency swap
	CurrencySwap struct {
		ID           uuid.UUID  `json:"id"`
		BusinessID   uuid.UUID  `json:"business_id"`
		FromAmount   Money      `json:"from"`
		ToAmount     Money      `json:"to"`
		ExchangeRate float64    `json:"exchangeRate" api:"rate"`
		Markup       Money      `json:"markup"`
		Status       string     `json:"status" api:"status"`
		FeeAmount    Money      `json:"fee"`
		CompletedAt  *time.Time `json:"completed_at"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at"`
	}

	// AllSwapsResponse schema for all currency swaps response
	AllSwapsResponse struct {
		Items []CurrencySwap `json:"items"`
		Page  PageInfo       `json:"page"`
	}
)
