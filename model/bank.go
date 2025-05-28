package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	// GenerateBankAccountRequest schema for generate bank account request
	GenerateBankAccountRequest struct {
		CustomerID                    string  `json:"customer_id"`
		Currency                      string  `json:"currency"`
		Reference                     string  `json:"reference"`
		BVN                           *string `json:"bvn,omitempty"`
		PhoneNumber                   *string `json:"phone_number,omitempty"`
		DocumentType                  *string `json:"document_type,omitempty"`
		Number                        *string `json:"document_number,omitempty"`
		IssuedCountryCode             *string `json:"issued_country_code,omitempty"`
		IssuedBy                      *string `json:"issued_by,omitempty"`
		IssuedDate                    *string `json:"issued_date,omitempty"`
		ExpirationDate                *string `json:"expiration_date,omitempty"`
		Country                       *string `json:"country,omitempty"`
		ZipCode                       *string `json:"zip_code,omitempty"`
		City                          *string `json:"city,omitempty"`
		Street                        *string `json:"street,omitempty"`
		State                         *string `json:"state,omitempty"`
		DateOfBirth                   *string `json:"date_of_birth,omitempty"`
		AgreementID                   *string `json:"agreement_id,omitempty"`
		DocumentFrontPage             *string `json:"document_front_page,omitempty"`
		DocumentBackPage              *string `json:"document_back_page,omitempty"`
		ProofOfAddressDoc             *string `json:"proof_of_address_doc,omitempty"`
		ActingAsIntermediary          *string `json:"acting_as_intermediary,omitempty"`    // true or false
		EmploymentStatus              *string `json:"employment_status,omitempty"`         // employed, homemaker, retired, self_employed, student, unemployed
		ExpectedMonthlyPayments       *string `json:"expected_monthly_payments,omitempty"` // "0_4999", "5000_9999", "10000_49999", "50000_plus"
		PrimaryPurpose                *string `json:"primary_purpose,omitempty"`           // business_transactions, charitable_donations, investment_purposes, payments_to_friends_or_family_abroad, personal_or_living_expenses, protect_wealth, purchase_goods_and_services, receive_payment_for_freelancing. other
		SourceOfFunds                 *string `json:"source_of_funds,omitempty"`           // "business_transactions", "charitable_donations", "investment_purposes", "payments_to_friends_or_family_abroad", "personal_or_living_expenses", "protect_wealth", "purchase_goods_and_services", "receive_payment_for_freelancing", "other"
		MostRecentOccupation          *string `json:"most_recent_occupation"`
		DocumentDescription           *string `json:"document_description"`
		AdditionalDocument            *string `json:"additional_document"` // passport/nin/other(with good enough description) for NGA
		AdditionalDocumentPurpose     *string `json:"additional_document_purpose"`
		AdditionalDocumentDescription *string `json:"additional_document_description"` // required if additional_document_purpose is "other"
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

	// TermsOfServiceResponse schema for terms of service response
	TermsOfServiceResponse struct {
		IsRequired bool   `json:"is_required"`
		URL        string `json:"url"`
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

	// NumberValidationResponse response for mobile number validation
	NumberValidationResponse struct {
		Phone               string `json:"phone"`
		Valid               bool   `json:"valid"`
		Mno                 string `json:"mno"`
		LocalFormat         string `json:"local_format"`
		InternationalFormat string `json:"international_format"`
		CountryPrefix       string `json:"country_prefix"`
		CountryCode         string `json:"country_code"`
		CountryName         string `json:"country_name"`
	}
)
