package model

import "time"

type (
	// BillerCategory is a bill payment category, e.g. airtime, electricity.
	BillerCategory struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// Biller is a billing entity configured under a bill payment category, e.g. MTN, DSTV.
	Biller struct {
		Code         string   `json:"code"`
		Name         string   `json:"name"`
		BillingTypes []string `json:"billing_types,omitempty"`
	}

	// BillerProduct is a payable product offered by a biller.
	BillerProduct struct {
		Code             string   `json:"code"`
		Name             string   `json:"name"`
		CategoryCode     string   `json:"category_code"`
		BillerCode       string   `json:"biller_code"`
		BillingType      string   `json:"billing_type,omitempty"`
		IsAmountEditable bool     `json:"is_amount_editable"`
		Amount           *float64 `json:"amount,omitempty"`
		MinAmount        *float64 `json:"min_amount,omitempty"`
		MaxAmount        *float64 `json:"max_amount,omitempty"`
	}

	// AllBillerProductsResponse schema for all biller products response
	AllBillerProductsResponse struct {
		Items []BillerProduct `json:"items"`
		Page  PageInfo        `json:"page"`
	}

	// ValidateBillerCustomerRequest is the request payload for validating a customer's
	// identifier (e.g. meter or smart card number) against a biller product before payment.
	ValidateBillerCustomerRequest struct {
		Code       string `json:"code"`
		CustomerID string `json:"customer_id"`
	}

	// ValidateBillerCustomerResponse contains the result of validating a customer's
	// identifier (e.g. meter or smart card number) against a biller product before payment.
	ValidateBillerCustomerResponse struct {
		CustomerName               string `json:"customer_name"`
		RequireValidationReference bool   `json:"require_validation_reference"`
		ValidationReference        string `json:"validation_reference,omitempty"`
	}

	// PayBillRequest is the request payload for initiating a bill payment.
	PayBillRequest struct {
		Code                string  `json:"code"`
		CustomerID          string  `json:"customer_id"`
		Amount              float64 `json:"amount"`
		ValidationReference *string `json:"validation_reference,omitempty"`
	}

	// BillPaymentMetadata holds provider-returned vend details, e.g. the prepaid meter token
	// and unit for an electricity bill payment.
	BillPaymentMetadata struct {
		Token *string `json:"token,omitempty"`
		Unit  *string `json:"unit,omitempty"`
	}

	// BillPaymentTransaction is a bill payment transaction.
	BillPaymentTransaction struct {
		ID                  string               `json:"id"`
		Code                string               `json:"code"`
		CustomerID          string               `json:"customer_id"`
		Amount              float64              `json:"amount"`
		Currency            string               `json:"currency"`
		ValidationReference *string              `json:"validation_reference,omitempty"`
		ProviderReference   *string              `json:"provider_reference,omitempty"`
		Status              string               `json:"status"`
		Metadata            *BillPaymentMetadata `json:"metadata,omitempty"`
		CreatedAt           time.Time            `json:"created_at"`
		UpdatedAt           *time.Time           `json:"updated_at,omitempty"`
	}
)
