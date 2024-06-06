package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

// InitiateDeposit makes request to Torus to initiate a deposit
func (c *Call) InitiateDeposit(ctx context.Context, request model.InitiateDepositRequest) (model.Deposit, error) {
	var (
		err       error
		response  model.Deposit
		path      = "v1/deposit"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// GetAllDeposits makes request to Torus to get all deposits
func (c *Call) GetAllDeposits(ctx context.Context, settled *bool) (model.DepositBatchResponse, error) {
	var (
		err      error
		response model.DepositBatchResponse
		params   = make(map[string]interface{})
		path     = "v1/deposits"
	)

	if settled != nil {
		params["settled"] = strconv.FormatBool(*settled)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetDepositID makes request to Torus to get deposit by its ID
func (c *Call) GetDepositID(ctx context.Context, id string) (model.Deposit, error) {
	var (
		err      error
		response model.Deposit
		path     = fmt.Sprintf("v1/deposit/%s", id)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// InternalFundsTransfer makes request to Torus to transfer funds between yield offerings
func (c *Call) InternalFundsTransfer(ctx context.Context, request model.FundTransferRequest) (model.Deposit, error) {
	var (
		err       error
		response  model.Deposit
		path      = "v1/transfer-funds"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// IntraTransfer makes request to Torus to transfer funds between two customers
func (c *Call) IntraTransfer(ctx context.Context, request model.IntraTransferRequest) (model.IntraTransferResponse, error) {
	var (
		err       error
		response  model.IntraTransferResponse
		path      = "v1/intra-transfer"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}
