package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ovalfi/go-sdk/model"
)

// GetKYCByCustomerID makes request to get KYC for a customer.
// It returns model.ErrCustomerKYCNotFound when the customer has no KYC record.
func (c *Call) GetKYCByCustomerID(ctx context.Context, customerID string) (model.KYCResponse, error) {
	var (
		err      error
		response model.KYCResponse
		path     = fmt.Sprintf("%s/%s", kycAPIVersion, customerID)
	)

	// the response carries the customer's identity and biometric details
	ctx = withRedactedLogging(ctx, path)
	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response.Data)
	// makeRequest only surfaces the error details as plain text, so map the known message to its sentinel
	if err != nil && err.Error() == model.ErrCustomerKYCNotFound.Error() {
		return response, model.ErrCustomerKYCNotFound
	}

	return response, err
}

// SubmitCustomerKYCDocument make request to submit a KYC document for a customer.
//
// Targets &response.Data, not &response, for the same reason as
// GetKYCByCustomerID above: makeRequest always decodes the already-unwrapped
// genericResponse.Data into whatever target is passed, never the outer
// envelope. Passing &response here previously meant every successful
// submission decoded the flat KYC object against KYCResponse{Status int,
// ...}, which fails on the object's own string "status" field.
func (c *Call) SubmitCustomerKYCDocument(
	ctx context.Context,
	customerID string,
	frontDocument *os.File,
	backDocument *os.File, // nil only if there is a front side
	documentType string,
	country string,
) (model.KYCResponse, error) {
	var (
		response model.KYCResponse
		formData = make(map[string]interface{})
		path     = fmt.Sprintf("%s/%s/document", kycAPIVersion, customerID)
	)

	// required front side
	formData["document_front_side"] = frontDocument

	// optional back side
	if backDocument != nil {
		formData["document_back_side"] = backDocument
	}

	// the rest text fields
	formData["document_type"] = documentType
	formData["country"] = country

	// the form carries the customer's identity documents and the response their KYC details
	ctx = withRedactedLogging(ctx, path)
	err := c.makeRequest(
		ctx,
		path,
		http.MethodPost,
		nil,
		nil,
		formData,
		nil,
		&response.Data,
	)
	return response, err
}

// VerifyCustomerKYC makes request to Torus to verify a customer kyc request
func (c *Call) VerifyCustomerKYC(ctx context.Context, customerID, idNumber, kycType string) (interface{}, error) {
	var (
		err      error
		response interface{}
		path     = fmt.Sprintf("%s/%s/%s/%s", kycAPIVersion, customerID, kycType, idNumber)
	)

	// the ID number (e.g. an SSN) is part of the path, so mask it in the logged path too
	ctx = withRedactedLogging(ctx, fmt.Sprintf("%s/%s/%s/***", kycAPIVersion, customerID, kycType))
	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, nil, &response)

	return response, err
}

// GetVerifyBiometricsLink makes request to get the link to verify biometrics
func (c *Call) GetVerifyBiometricsLink(ctx context.Context, customerID string) (string, error) {
	var (
		err      error
		response string
		path     = fmt.Sprintf("%s/%s/biometrics", kycAPIVersion, customerID)
	)

	// the response is a one-time verification link for the customer
	ctx = withRedactedLogging(ctx, path)
	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err

}

// GetVerifyCustomerKYC makes request to get the link to verify biometrics.
func (c *Call) GetVerifyCustomerKYC(ctx context.Context, customerID string, country, hasExpiredID *string) (model.VerifyCustomerKYCResponse, error) {
	return c.getVerifyCustomerKYC(ctx, customerID, country, hasExpiredID, nil, nil)
}

// GetVerifyCustomerKYCWithSessionType is GetVerifyCustomerKYC plus preferredSessionType
// ("redirect_url" | "access_token")
func (c *Call) GetVerifyCustomerKYCWithSessionType(ctx context.Context, customerID string, country, hasExpiredID *string, preferredSessionType *model.SessionType) (model.VerifyCustomerKYCResponse, error) {
	return c.getVerifyCustomerKYC(ctx, customerID, country, hasExpiredID, preferredSessionType, nil)
}

// GetProofOfAddressVerificationLink gets a session for collecting only a proof of address document
func (c *Call) GetProofOfAddressVerificationLink(ctx context.Context, customerID string, country *string, preferredSessionType *model.SessionType) (model.VerifyCustomerKYCResponse, error) {
	purpose := model.VerificationPurposeProofOfAddress
	return c.getVerifyCustomerKYC(ctx, customerID, country, nil, preferredSessionType, &purpose)
}

func (c *Call) getVerifyCustomerKYC(ctx context.Context, customerID string, country, hasExpiredID *string, preferredSessionType *model.SessionType, verificationPurpose *model.VerificationPurpose) (model.VerifyCustomerKYCResponse, error) {
	var (
		err      error
		response model.VerifyCustomerKYCResponse
		params   = make(map[string]interface{})
		path     = fmt.Sprintf("%s/%s/verify", kycAPIVersion, customerID)
	)

	if country != nil {
		params["country"] = *country
	}

	if hasExpiredID != nil {
		params["has_expired_id"] = *hasExpiredID
	}
	if preferredSessionType != nil {
		params["preferred_session_type"] = string(*preferredSessionType)
	}
	if verificationPurpose != nil {
		params["verification_purpose"] = string(*verificationPurpose)
	}
	// the response is a one-time verification link for the customer
	ctx = withRedactedLogging(ctx, path)
	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err

}
