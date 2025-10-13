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
	CreatedAt                      time.Time   `json:"created_at"`
	UpdatedAt                      time.Time   `json:"updated_at"`
	DeletedAt                      *time.Time  `json:"deleted_at"`
	Documents                      []Document  `json:"documents"`
}

type Document struct {
	ID               string      `json:"id"`
	BusinessID       string      `json:"businessId"`
	CustomerID       string      `json:"customerId"`
	CustomerKYCID    string      `json:"customerKycId"`
	DocType          string      `json:"docType"`
	DocSubtype       string      `json:"docSubtype"`
	Description      interface{} `json:"description"`
	Status           string      `json:"status"`
	FailureNotes     interface{} `json:"failureNotes"`
	Extension        string      `json:"extension"`
	Label            string      `json:"label"`
	IsIdentity       bool        `json:"isIdentity"`
	IsProofOfAddress bool        `json:"isProofOfAddress"`
	ProviderPayload  interface{} `json:"providerPayload"`
	CreatedAt        string      `json:"createdAt"`
	UpdatedAt        interface{} `json:"updatedAt"`
	VerifiedAt       interface{} `json:"verifiedAt"`
	DeletedAt        interface{} `json:"deletedAt"`
}

type VerifyCustomerKYCResponse struct {
	BusinessID  string `json:"businessID"`
	URL         string `json:"url"`
	CustomerID  string `json:"customerID"`
	KYCProvider string `json:"kycProvider"`
}

type VerifyCustomerKYCRequest struct {
	Country *string `json:"country,omitempty"`
}
