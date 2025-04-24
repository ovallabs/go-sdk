package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type (
	// InitiateCardRequest to initiate card request payload
	InitiateCardRequest struct {
		CustomerID  uuid.UUID `json:"customer_id"  validate:"required"`
		Reference   string    `json:"reference"  validate:"required"`
		DateOfBirth string    `json:"date_of_birth" validate:"required"` // format: DD-MMM-YYYY (17-JAN-1985)
		SSN         string    `json:"ssn" validate:"required,len=4"`     // Social Security Number of the user (format: Last four ####)
		Phone       string    `json:"phone" validate:"required"`         // Phone number of the user (format: +15557771234)
		Address     string    `json:"address" validate:"required"`       // Address line of the user (PO Boxes are not allowed)
		City        string    `json:"city" validate:"required"`          // City of the user
		State       string    `json:"state" validate:"required,len=2"`   // State of the user
		PostalCode  string    `json:"postal_code" validate:"required"`   // Postal code of the user
		IPAddress   string    `json:"ip_address" validate:"required"`    // IP address of the user
		RedirectURI string    `json:"redirect_uri" validate:"required"`
	}

	// CompleteCardRequest to complete card request payload
	CompleteCardRequest struct {
		CustomerID  uuid.UUID `json:"customer_id"  validate:"required"`
		AuthCode    string    `json:"auth_code"  validate:"required"`
		RedirectURI string    `json:"redirect_uri" validate:"required"`
	}

	// GetLinkToAddCardReq to get link to add payment card
	GetLinkToAddCardReq struct {
		CustomerID  uuid.UUID `json:"customer_id" validate:"required"`
		RedirectURI string    `json:"redirect_uri" validate:"required"`
		Phone       *string   `json:"phone"`
		DirectDebit *bool     `json:"direct_debit"`
	}

	// PaymentCard schema represents entity that contains all needed information of a customer payment card
	PaymentCard struct {
		ID             uuid.UUID       `json:"id"`
		BusinessID     uuid.UUID       `json:"business_id"`
		CustomerID     uuid.UUID       `json:"customer_id"`
		FirstName      string          `json:"first_name"`
		LastName       string          `json:"last_name"`
		CardBrand      string          `json:"card_brand"`
		FirstSixDigits string          `json:"first_six_digits"`
		LastFourDigits string          `json:"last_four_digits"`
		ExpiryDate     string          `json:"expiry_date"`
		Type           PaymentCardType `json:"type"`
		IssuerName     string          `json:"issuer_name"`
		Status         string          `json:"status"`
		BillingAddress postgres.Jsonb  `json:"billing_address"`
		CreatedAt      time.Time       `json:"created_at"`
	}

	// AllPaymentCardsResponse schema for all payment cards response
	AllPaymentCardsResponse struct {
		Items *[]PaymentCard `json:"items"`
		Page  PageInfo       `json:"page"`
	}

	// PaymentCardType a string representation of customer payment card type
	PaymentCardType string
)

const (
	// CardTypeDebit represents debit card type
	CardTypeDebit PaymentCardType = "debit"
	// CardTypePrepaid represents prepaid card type
	CardTypePrepaid PaymentCardType = "prepaid"
	// CardTypeCredit represents prepaid card type
	CardTypeCredit PaymentCardType = "credit"
	// CardTypeUnknown represents unknown card type
	CardTypeUnknown PaymentCardType = "unknown"
)
