package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const accountRequestAPIVersion = "v1/payments/account-request"

// GetAccountOpeningRequests makes a request to Torus to get all account opening requests
func (c *Call) GetAccountOpeningRequests(ctx context.Context, customerID, status, currency string, opened *bool, page model.Page) (model.AllAccountOpeningRequests, error) {
	var (
		err      error
		response model.AllAccountOpeningRequests
		params   = make(map[string]interface{})
		path     = accountRequestAPIVersion
	)

	if customerID != "" {
		params["customer_id"] = customerID
	}
	if status != "" {
		params["status"] = status
	}
	if currency != "" {
		params["currency"] = currency
	}
	if opened != nil {
		params["opened"] = strconv.FormatBool(*opened)
	}

	if page != (model.Page{}) {
		helpers.FillParamsWithPage(params, page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetAccountOpeningRequestByID makes a request to Torus to get account opening request by its ID
func (c *Call) GetAccountOpeningRequestByID(ctx context.Context, accountID string) (model.AccountOpeningRequest, error) {
	var (
		err      error
		response model.AccountOpeningRequest
		path     = fmt.Sprintf("%s/%s", accountRequestAPIVersion, accountID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
