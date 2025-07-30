package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ovalfi/go-sdk/model"
)

// GetKYCByCustomerID makes request to get KYC for a customer
func (c *Call) GetKYCByCustomerID(ctx context.Context, customerID string) (model.KYCResponse, error) {
	var (
		err      error
		response model.KYCResponse
		path     = fmt.Sprintf("%s/%s", kycAPIVersion, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// SubmitCustomerKYCDocument allows you to submit a KYC document for a customer
func (c *Call) SubmitCustomerKYCDocument(
	ctx context.Context,
	customerID string,
	frontDocument *os.File,
	backDocument *os.File, // nil if you only have a front side
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

	// makeRequest
	err := c.makeRequest(
		ctx,
		path,
		http.MethodPost,
		nil,      // signature
		nil,      // url params
		formData, // form data
		nil,
		&response,
	)
	return response, err
}

// VerifyCustomerKYC makes request to Torus to verify a customer kyc request
func (c *Call) VerifyCustomerKYC(ctx context.Context, customerID, idNumber, kycType string) error {
	var (
		err  error
		path = fmt.Sprintf("%s/%s/%s/%s", kycAPIVersion, customerID, kycType, idNumber)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, nil, nil)

	return err
}
