package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const transactionAPIVersion = "v1/transaction"

// GetTransactions makes request to Torus to get all transactions
func (c *Call) GetTransactions(ctx context.Context, customerID, yieldOfferingID, status, reference, batchDate string, amount *float64, dateBetween model.DateBetween, page model.Page) (model.AllTransactionsResponse, error) {
	var (
		err      error
		response model.AllTransactionsResponse
		params   = make(map[string]interface{})
		path     = transactionAPIVersion
	)

	if customerID != "" {
		params["customer_id"] = customerID
	}
	if yieldOfferingID != "" {
		params["yield_offering_id"] = yieldOfferingID
	}
	if status != "" {
		params["status"] = status
	}
	if reference != "" {
		params["reference"] = reference
	}
	if batchDate != "" {
		params["batch_date"] = batchDate
	}
	if amount != nil {
		params["amount"] = strconv.FormatFloat(*amount, 'f', -1, 64)
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

// CancelTransaction makes request to Torus to cancel transaction
func (c *Call) CancelTransaction(ctx context.Context, transactionID, transactionType, reason string) error {
	var (
		err    error
		params = map[string]interface{}{"type": transactionType, "reason": reason}
		path   = fmt.Sprintf("%s/%s", transactionAPIVersion, transactionID)
	)

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, params, nil, nil, nil)

	return err
}

// CancelBatchTransaction makes request to Torus to cancel batch transaction
func (c *Call) CancelBatchTransaction(ctx context.Context, batchDate, transactionType, currency, reason string) error {
	var (
		err    error
		params = map[string]interface{}{"type": transactionType, "reason": reason}
		path   = fmt.Sprintf("v1/batch/%s", batchDate)
	)

	if currency != "" {
		params["currency"] = currency
	}

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, params, nil, nil, nil)

	return err
}

// GetBalances makes request to Torus to get business balances
func (c *Call) GetBalances(ctx context.Context) (map[string]float64, error) {
	var (
		err      error
		response map[string]float64
		path     = "v1/balances"
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
