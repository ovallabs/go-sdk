package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const bankAPIVersion = "v1/payments/banks"

// ResolveBankAccount makes a request to Torus to resolve bank account
func (c *Call) ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetails, error) {
	var (
		err      error
		response model.AccountDetails
		path     = fmt.Sprintf("%s/resolve-account", bankAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetBanks makes request to Torus to get list of banks
func (c *Call) GetBanks(ctx context.Context) ([]model.BankCode, error) {
	var (
		err      error
		response []model.BankCode
		path     = bankAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GenerateBankAccount makes request to Torus to generate bank account
func (c *Call) GenerateBankAccount(ctx context.Context, request model.GenerateBankAccountRequest) (model.BankAccount, error) {
	var (
		err       error
		response  model.BankAccount
		path      = fmt.Sprintf("%s/account", bankAPIVersion)
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// GetBankAccount makes request to Torus to get bank account
func (c *Call) GetBankAccount(ctx context.Context, customerID, currency string) (model.BankAccount, error) {
	var (
		err      error
		response model.BankAccount
		params   = map[string]interface{}{
			"customer_id": customerID,
			"currency":    currency,
		}
		path = fmt.Sprintf("%s/account", bankAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}
