package api

import (
	"context"
	"errors"
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

// GetDepositByIDOrReference makes a request to Torus to get a deposit by its ID or by its Reference
func (c *Call) GetDepositByIDOrReference(ctx context.Context, id, reference *string) (model.Deposit, error) {
	var (
		err      error
		response model.Deposit
		basePath = "v1/deposit/search"
		query    string
		fullPath string
	)

	if id == nil && reference == nil {
		return model.Deposit{}, errors.New("must provide either 'id' or 'reference'")
	}

	idVal := ""
	if id != nil {
		idVal = *id
	}

	refVal := ""
	if reference != nil {
		refVal = *reference
	}

	// 2. Core logic using the safely dereferenced values (idVal, refVal)
	if idVal != "" && refVal == "" {
		query = fmt.Sprintf("?id=%s", idVal)
	} else if refVal != "" && idVal == "" {
		query = fmt.Sprintf("?reference=%s", refVal)
	} else if idVal != "" && refVal != "" {
		return model.Deposit{}, errors.New("cannot query deposit with both 'id' and 'reference'. Provide only one")
	} else {
		return model.Deposit{}, errors.New("must provide either 'id' or 'reference'")
	}

	fullPath = basePath + query

	err = c.makeRequest(ctx, fullPath, http.MethodGet, nil, nil, nil, nil, &response)

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
