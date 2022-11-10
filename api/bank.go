package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const (
	bankEndpoint        = "v1/payments/banks"
	bankAccountEndpoint = "v1/payments/banks/account"
)

// ResolveBankAccount makes an API request using Call to resolve bank account
func (c Call) ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetailResponse, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, bankEndpoint, "/resolve-account")

	response := struct {
		Data model.AccountDetailResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.AccountDetailResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.AccountDetailResponse{}, model.ErrNetworkError
	}

	return response.Data, nil
}

// GetBanks makes an API request using Call to get list of banks
func (c Call) GetBanks(ctx context.Context) ([]model.BankCodeResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, bankEndpoint)

	response := struct {
		Data []model.BankCodeResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return response.Data, err
	}

	if res.StatusCode() != http.StatusOK {
		return response.Data, model.ErrNetworkError
	}

	return response.Data, nil
}

// GenerateBankAccount makes an API request to generate bank account
func (c *Call) GenerateBankAccount(ctx context.Context, request model.BankAccountRequest) (model.BankAccountResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, bankAccountEndpoint)

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.BankAccountResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.BankAccountResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.BankAccountResponse{}, model.ErrNetworkError
	}

	return response.Data, nil
}

// GetBankAccount makes an API request to get bank account
func (c *Call) GetBankAccount(ctx context.Context, customerID uuid.UUID) (model.BankAccountResponse, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, bankAccountEndpoint, customerID)
	response := struct {
		Data model.BankAccountResponse `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return model.BankAccountResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.BankAccountResponse{}, model.ErrNetworkError
	}

	return response.Data, nil
}
