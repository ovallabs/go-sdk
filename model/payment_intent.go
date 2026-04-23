package model

import "github.com/google/uuid"

type (
	// CreateCustomerPaymentIntentRequest struct to create payment intent request
	CreateCustomerPaymentIntentRequest struct {
		CustomerID  string  `json:"customer_id"`
		Amount      float64 `json:"amount"`
		Currency    string  `json:"currency"`
		Reference   *string `json:"reference,omitempty"`
		RedirectURL *string `json:"redirect_url,omitempty"`
		PhoneNumber *string `json:"phone_number,omitempty"`
	}

	// CompleteCustomerPaymentIntentRequest struct to complete payment intent
	CompleteCustomerPaymentIntentRequest struct {
		PaymentIntentID string       `json:"payment_intent_id"`
		CustomerID      string       `json:"customer_id"`
		PaymentMethod   string       `json:"payment_method"`
		Country         string       `json:"country"`
		Provider        string       `json:"provider"`
		MobileMoney     *MobileMoney `json:"mobile_money,omitempty"`
		Redirection     *Redirection `json:"redirection,omitempty"`
	}

	// AuthenticateCustomerPaymentIntentRequest struct to authenticate payment intent
	AuthenticateCustomerPaymentIntentRequest struct {
		PaymentIntentID  string `json:"payment_intent_id"`
		CustomerID       string `json:"customer_id"`
		ConfirmationCode string `json:"confirmation_code"`
	}

	// CreateCustomerPaymentIntentResponse struct for create payment intent response
	CreateCustomerPaymentIntentResponse struct {
		CustomerPaymentIntent CustomerPaymentIntent `json:"customerPaymentIntent"`
	}

	// CustomerPaymentIntent payment intent response object
	CustomerPaymentIntent struct {
		ID                uuid.UUID `json:"id"`
		CustomerID        uuid.UUID `json:"customer_id"`
		BusinessID        uuid.UUID `json:"business_id"`
		ProviderReference string    `json:"provider_reference,omitempty"`
		Amount            Money     `json:"amount"`
		PaymentMethod     string    `json:"payment_method,omitempty"`
		Status            string    `json:"status"`
		Country           *string   `json:"country,omitempty"`
		PaymentURL        *string   `json:"payment_url,omitempty"`
	}

	// MobileMoney mobile money details object
	MobileMoney struct {
		MSISDN                  string  `json:"msisdn"`
		OTP                     *string `json:"otp"`
		OnSuccessRedirectionURL *string `json:"on_success_redirection_url"`
		OnFailedRedirectionURL  *string `json:"on_failed_redirection_url"`
		OnCancelRedirectionURL  *string `json:"on_cancel_redirection_url"`
		OnFinishRedirectionURL  *string `json:"on_finish_redirection_url"`
		Workflow                *string `json:"workflow"`
	}

	// Redirection generic redirection details object
	Redirection struct {
		OnSuccessRedirectionURL *string `json:"on_success_redirection_url"`
		OnFailedRedirectionURL  *string `json:"on_failed_redirection_url"`
		OnCancelRedirectionURL  *string `json:"on_cancel_redirection_url"`
		OnFinishRedirectionURL  *string `json:"on_finish_redirection_url"`
	}
)
