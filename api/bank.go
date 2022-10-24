package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

const paymentEndpoint = "v1/payments"

// ResolveBankAccount makes an API request using Call to resolve bank account
func (c Call) ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetailResponse, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, paymentEndpoint, "/resolve-account")

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
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, paymentEndpoint, "/banks")

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
