package api

import (
	"context"
	"fmt"
	"github.com/ovalfi/go-sdk/model"
	"net/http"
	"os"
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

func (c *Call) SubmitCustomerKYCDocument(
	ctx context.Context,
	customerID string,
	document *os.File,
	documentType string,
	country string,
) (model.KYCResponse, error) {
	var (
		err      error
		response model.KYCResponse
		formData = make(map[string]interface{})
		path     = fmt.Sprintf("%s/%s/document-verification", kycAPIVersion, customerID)
	)

	// Prepare form data
	formData["document"] = document
	formData["documentType"] = documentType
	formData["country"] = country

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, formData, nil, &response)

	return response, err
}
