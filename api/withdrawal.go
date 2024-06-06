package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const withdrawalAPIVersion = "v1/withdrawal"

// InitiateWithdrawal makes request to Torus to initiate a withdrawal
func (c *Call) InitiateWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	var (
		err       error
		response  model.Withdrawal
		path      = withdrawalAPIVersion
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// FiatWithdrawal makes request to Torus to withdraw to a provided bank account
func (c *Call) FiatWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	var (
		err       error
		response  model.Withdrawal
		path      = fmt.Sprintf("%s/fiat", withdrawalAPIVersion)
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// CryptoWithdrawal makes request to Torus to withdraw to a specified crypto wallet address
func (c *Call) CryptoWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	var (
		err       error
		response  model.Withdrawal
		path      = fmt.Sprintf("%s/crypto", withdrawalAPIVersion)
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// FeeWithdrawal makes request to Torus for fee withdrawal
func (c *Call) FeeWithdrawal(ctx context.Context, request model.FeeWithdrawalRequest) (model.FeeWithdrawalResponse, error) {
	var (
		err       error
		response  model.FeeWithdrawalResponse
		path      = fmt.Sprintf("%s/fee", withdrawalAPIVersion)
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}
