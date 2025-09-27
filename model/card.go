package model

import "github.com/google/uuid"

type (
	// CreateCustomerCardRequest schema
	CreateCustomerCardRequest struct {
		CustomerID string `json:"customer_id"`
		CardType   string `json:"card_type"`
		ID         struct {
			Type     string `json:"type"`
			Value    string `json:"value"`
			Country  string `json:"country"`
			ImageURL string `json:"image_url"`
		} `json:"id"`
		Reference     string `json:"reference"`
		PreferredName string `json:"preferred_name"`
		Address       string `json:"address"`
		City          string `json:"city"`
		Country       string `json:"country"`
		IDNumber      string `json:"id_number"`
		StateRegion   string `json:"state_region"`
		PostalCode    string `json:"postal_code"`
		BirthDate     string `json:"birth_date"`
		Phone         string `json:"phone"`
	}

	// CreateCustomerCardRequestV2 schema
	CreateCustomerCardRequestV2 struct {
		CardType      string `json:"card_type"`
		ImageURL      string `json:"image_url"`
		PreferredName string `json:"preferred_name"`
		GenerateBankAccountRequest
	}

	// FreezeCardRequest schema
	FreezeCardRequest struct {
		CardID       string `json:"card_id"`
		CustomerID   string `json:"customer_id"`
		FreezeCard   string `json:"freeze_card"`
		FreezeReason string `json:"freeze_reason"`
	}

	// Card schema
	Card struct {
		ID             uuid.UUID      `json:"id"`
		BusinessID     uuid.UUID      `json:"business_id"`
		CustomerID     uuid.UUID      `json:"customer_id"`
		CardName       string         `json:"card_name"`
		LastFourDigits string         `json:"last_four_digits"`
		FirstSixDigits string         `json:"first_six_digits"`
		ExpiryDate     string         `json:"expiry_date"`
		Frozen         bool           `json:"frozen"`
		IssuerName     string         `json:"issuer_name"`
		Type           string         `json:"type"`
		Status         string         `json:"status"`
		BillingAddress BillingAddress `json:"billing_address"`
		IssuedAt       interface{}    `json:"issued_at"`
		CreatedAt      interface{}    `json:"created_at"`
	}

	// BillingAddress schema
	BillingAddress struct {
		City        string `json:"city"`
		Address     string `json:"address"`
		Country     string `json:"country"`
		PostalCode  string `json:"postal_code"`
		StateRegion string `json:"state_region"`
	}

	// AllCardsResponse schema for all payment cards response
	AllCardsResponse struct {
		Items *[]Card  `json:"items"`
		Page  PageInfo `json:"page"`
	}

	// FundCustomerCardRequest schema
	FundCustomerCardRequest struct {
		CardID            string          `json:"card_id"`
		CustomerID        string          `json:"customer_id"`
		TransferAmount    float64         `json:"transfer_amount"`
		TransferNarration string          `json:"transfer_narration"`
		TransactionFlow   TransactionFlow `json:"transaction_flow"`
	}

	// VaultedCardDetails secure vaulted card details
	VaultedCardDetails struct {
		FullPAN    string `json:"full_pan"`     // full card number
		CVV        string `json:"cvv"`          // sensitive
		ExpiryDate string `json:"expiry_date"`  // e.g., "09"
		NameOnCard string `json:"name_on_card"` // optional
		Issuer     string `json:"issuer"`       // optional, duplicate allowed
	}

	// CustomerPaymentSessionRequest schema for customer payment session request
	CustomerPaymentSessionRequest struct {
		CustomerID    string  `json:"customer_id" validate:"required"`
		PaymentMethod *string `json:"payment_method,omitempty"`
		Amount        float64 `json:"amount" validate:"required"`
		Currency      string  `json:"currency" validate:"required"`
		Reference     string  `json:"reference" validate:"required"`
		FailureURL    string  `json:"failure_url" validate:"required"`
		SuccessURL    string  `json:"success_url" validate:"required"`
		DisplayName   string  `json:"display_name" validate:"required"`
		Country       string  `json:"country" validate:"required,max=2"`
		Remarks       *string `json:"remarks"`
	}

	// CustomerPaymentSessionResponse schema for customer payment session response
	CustomerPaymentSessionResponse struct {
		Reference     string `json:"reference"`
		CustomerID    string `json:"customer_id"`
		SessionID     string `json:"session_id"`
		SessionSecret string `json:"session_secret"`
		SessionToken  string `json:"session_token"`
	}
)
