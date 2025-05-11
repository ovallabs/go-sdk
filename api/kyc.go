package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

const kycAPIVersion = "v1/kycs"

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
