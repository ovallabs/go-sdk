// Package model/kyc.go
package model

import (
	"time"
)

type KYCResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    KYCData     `json:"data"`
}

type KYCData struct {
	ID                             string      `json:"id"`
	BusinessID                     string      `json:"business_id"`
	CustomerID                     string      `json:"customer_id"`
	KYCProvider                    string      `json:"kyc_provider"`
	KYCType                        string      `json:"kyc_type"`
	ProviderContactID              string      `json:"provider_contact_id"`
	Name                           string      `json:"name"`
	Sex                            string      `json:"sex"`
	MaritalStatus                  string      `json:"marital_status"`
	DateOfBirth                    string      `json:"date_of_birth"`
	Email                          string      `json:"email"`
	PhoneNumber                    string      `json:"phone_number"`
	Country                        string      `json:"country"`
	ContactType                    string      `json:"contact_type"`
	Status                         string      `json:"status"`
	Identity                       string      `json:"identity"`
	IdentityType                   string      `json:"identity_type"`
	IdentityConfirmed              bool        `json:"identity_confirmed"`
	IdentityVerificationStatus     string      `json:"identity_verification_status"`
	IdentityDocumentVerified       bool        `json:"identity_document_verified"`
	ProofOfAddressDocumentVerified bool        `json:"proof_of_address_document_verified"`
	TaxIDNumber                    string      `json:"tax_id_number"`
	TaxCountry                     string      `json:"tax_country"`
	TaxState                       string      `json:"tax_state"`
	TaxIDVerified                  bool        `json:"tax_id_verified"`
	TaxVerificationStatus          bool        `json:"tax_verification_status"`
	AMLDetails                     interface{} `json:"aml_details"`
	BiometricsVerified             bool        `json:"biometrics_verified"`
	BiometricsVerificationStatus   interface{} `json:"biometrics_verification_status"`
	BiometricsDocumentDetails      interface{} `json:"biometrics_document_details"`
	HasBiometricsVerification      bool        `json:"has_biometrics_verification"`
	HasTaxVerification             bool        `json:"has_tax_verification"`
	HasIdentityVerification        bool        `json:"has_identity_verification"`
	CreatedAt                      time.Time   `json:"created_at"`
	UpdatedAt                      time.Time   `json:"updated_at"`
	DeletedAt                      *time.Time  `json:"deleted_at"`
	Documents                      []Document  `json:"documents"`
}

type Document struct {
	ID               string      `json:"id"`
	BusinessID       string      `json:"business_id"`
	CustomerID       string      `json:"customer_id"`
	CustomerKYCID    string      `json:"customer_kyc_id"`
	DocType          string      `json:"doc_type"`
	DocSubtype       string      `json:"doc_subtype"`
	Description      interface{} `json:"description"`
	Status           string      `json:"status"`
	FailureNotes     interface{} `json:"failure_notes"`
	Extension        string      `json:"extension"`
	FrontSideLabel   string      `json:"front_side_label"`
	BackSideLabel    string      `json:"back_side_label"`
	IsIdentity       bool        `json:"is_identity"`
	IsProofOfAddress bool        `json:"is_proof_of_address"`
	ProviderPayload  interface{} `json:"provider_payload"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        interface{} `json:"updated_at"`
	VerifiedAt       interface{} `json:"verified_at"`
	DeletedAt        interface{} `json:"deleted_at"`
}

type VerifyCustomerKYCResponse struct {
	BusinessID  string `json:"business_id"`
	URL         string `json:"url"`
	CustomerID  string `json:"customer_id"`
	KYCProvider string `json:"kyc_provider"`
	SessionType string `json:"session_type,omitempty"`
	FlowID      string `json:"flow_id,omitempty"`
}

type VerifyCustomerKYCRequest struct {
	Country *string `json:"country,omitempty"`
}
