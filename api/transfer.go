package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const (
	transferAPIVersion = "v1/transfers"

	customerTransferAPIVersion = "v1/customer-transfers"

	settlementAPIVersion = "v1/settlement"
)

// InitiateTransfer makes request to Torus to initiate transfer
func (c *Call) InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.TransferResponse, error) {
	var (
		err       error
		response  model.TransferResponse
		path      = customerTransferAPIVersion
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// GetExchangeRates makes request to Torus to get exchange rate
func (c *Call) GetExchangeRates(ctx context.Context, amount float64, sourceCurrency, destinationCurrency string) (model.ExchangeRateDetails, error) {
	var (
		err      error
		response model.ExchangeRateDetails
		params   = map[string]interface{}{
			"source_currency":      sourceCurrency,
			"destination_currency": destinationCurrency,
		}
		path = fmt.Sprintf("%s/quote", customerTransferAPIVersion)
	)

	strAmount := strconv.FormatFloat(amount, 'f', -1, 64)
	params["amount"] = strAmount

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetTransferByID makes request to Torus to get transfer by its ID
func (c *Call) GetTransferByID(ctx context.Context, transferID string) (model.Transfer, error) {
	var (
		err      error
		response model.Transfer
		path     = fmt.Sprintf("%s/%s", customerTransferAPIVersion, transferID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// DeleteTransfer makes request to Torus to delete transfer by its ID
func (c *Call) DeleteTransfer(ctx context.Context, transferID, reason string) error {
	var (
		err    error
		params = map[string]interface{}{
			"reason": reason,
		}
		path = fmt.Sprintf("%s/%s", customerTransferAPIVersion, transferID)
	)

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, params, nil, nil, nil)

	return err
}

// DeleteTransferBatch makes request to Torus to delete transfer by batch date
func (c *Call) DeleteTransferBatch(ctx context.Context, batchDate, currency, reason string) error {
	var (
		err    error
		params = map[string]interface{}{
			"currency": currency,
			"reason":   reason,
		}
		path = fmt.Sprintf("%s/delete-by-batch/%s", customerTransferAPIVersion, batchDate)
	)

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, params, nil, nil, nil)

	return err
}

// InitiateTerminalTransfer makes request to Torus to initiate terminal transfer
func (c *Call) InitiateTerminalTransfer(ctx context.Context, request model.InitiateTerminalTransferRequest) (model.TerminalTransfer, error) {
	var (
		err      error
		response model.TerminalTransfer
		path     = transferAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetTerminalTransfers makes request to Torus to get all terminal transfers
func (c *Call) GetTerminalTransfers(ctx context.Context, status, sourceCurrency, destinationCurrency string, dateBetween model.DateBetween, page model.Page) (model.AllTransfersResponse, error) {
	var (
		err      error
		response model.AllTransfersResponse
		params   = make(map[string]interface{})
		path     = transferAPIVersion
	)

	if status != "" {
		params["status"] = status
	}
	if sourceCurrency != "" {
		params["source_currency"] = sourceCurrency
	}
	if destinationCurrency != "" {
		params["destination_currency"] = destinationCurrency
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

// GetTerminalTransferByID makes request to Torus to get terminal transfer by its ID
func (c *Call) GetTerminalTransferByID(ctx context.Context, transferID string) (model.TerminalTransfer, error) {
	var (
		err      error
		response model.TerminalTransfer
		path     = fmt.Sprintf("%s/%s", transferAPIVersion, transferID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetSettlementByID makes request to Torus to get settlement by its ID
func (c *Call) GetSettlementByID(ctx context.Context, settlementID string) (model.Settlement, error) {
	var (
		err      error
		response model.Settlement
		path     = fmt.Sprintf("%s/%s", settlementAPIVersion, settlementID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
