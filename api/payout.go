package api

import (
	"context"
	"fmt"
	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
	"net/http"
	"os"
)

const payoutAPIVersion = "v1/payouts"

// GetPayoutByID makes a request to Torus to get the payout by its ID.
func (c *Call) GetPayoutByID(ctx context.Context, payoutID string) (model.PayoutResponse, error) {
	var (
		err      error
		response model.PayoutResponse
		path     = fmt.Sprintf("%s/%s", payoutAPIVersion, payoutID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// InitiateDirectBulkPayout makes a request to Torus to initiate a bulk payout
func (c *Call) InitiateDirectBulkPayout(ctx context.Context, request model.InitiateBulkPayoutRequest) (model.PayoutDetails, error) {
	var (
		err      error
		response model.PayoutDetails
		path     = payoutAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// InitiatePayout makes a request to Torus to initiate a bulk payout
func (c *Call) InitiatePayout(ctx context.Context, currency, payoutType, beneficiaryType, remarks string, document *os.File) (model.PayoutDetails, error) {
	var (
		err      error
		response model.PayoutDetails
		formData = make(map[string]interface{})
		path     = fmt.Sprintf("%s/upload", payoutAPIVersion)
	)

	formData["currency"] = currency
	formData["payout_type"] = payoutType
	formData["beneficiary_type"] = beneficiaryType
	formData["remarks"] = remarks
	formData["document"] = document

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, formData, nil, &response)

	return response, err
}

// GetAllPayouts makes request to Torus to get all payouts
func (c *Call) GetAllPayouts(ctx context.Context, status, search string, dateBetween model.DateBetween, page model.Page) (model.AllPayoutsResponse, error) {
	var (
		err      error
		response model.AllPayoutsResponse
		params   = make(map[string]interface{})
		path     = payoutAPIVersion
	)

	if status != "" {
		params["status"] = status
	}
	if search != "" {
		params["search"] = search
	}
	if dateBetween != (model.DateBetween{}) {
		helpers.FillParamsWithDateInterval(params, dateBetween)
	}
	if page != (model.Page{}) {
		helpers.FillParamsWithPage(params, page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// CancelPayout makes request to Torus to cancel payout
func (c *Call) CancelPayout(ctx context.Context, request model.CancelPayoutRequest) error {
	var (
		err  error
		path = fmt.Sprintf("%s/cancel", payoutAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, nil)

	return err
}

// UpdatePayoutAccount makes request to Torus to update payout account by its ID
func (c *Call) UpdatePayoutAccount(ctx context.Context, payoutID string, request model.TransferBeneficiaryDetails) error {
	var (
		err  error
		path = fmt.Sprintf("%s/accounts/%s", payoutAPIVersion, payoutID)
	)

	err = c.makeRequest(ctx, path, http.MethodPut, nil, nil, nil, request, nil)

	return err
}

// GetPayoutConfig makes request to Torus to get payout config
func (c *Call) GetPayoutConfig(ctx context.Context, currency string) (model.BulkPayoutConfig, error) {
	var (
		err      error
		response model.BulkPayoutConfig
		path     = fmt.Sprintf("%s/config/%s", payoutAPIVersion, currency)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetPayoutDocumentTemplate makes request to Torus to get payout document template
func (c *Call) GetPayoutDocumentTemplate(ctx context.Context, currency, docType string) (string, error) {
	var (
		err      error
		response string
		params   = make(map[string]interface{})
		path     = fmt.Sprintf("%s/template", payoutAPIVersion)
	)

	if currency != "" {
		params["currency"] = currency
	}
	if docType != "" {
		params["type"] = docType
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}
