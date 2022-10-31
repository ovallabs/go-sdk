package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const withdrawalAPIVersion = "v1/withdrawal"

// InitiateWithdrawal makes an API request using Call to initiate a withdrawal
func (c *Call) InitiateWithdrawal(ctx context.Context, request model.InitiateWithdrawalRequest) (model.Withdrawal, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, withdrawalAPIVersion)

	fL := c.logger.With().Str("func", "InitiateWithdrawal").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.Withdrawal{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// FiatWithdrawal makes an API request to withdrawal to a provided bank account
func (c *Call) FiatWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, withdrawalAPIVersion, "/fiat")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.Withdrawal{}, model.ErrNetworkError
	}

	return response.Data, nil
}

// CryptoWithdrawal makes an API request to withdrawal to a specified crypto wallet address
func (c Call) CryptoWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, withdrawalAPIVersion, "/crypto")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.Withdrawal{}, model.ErrNetworkError
	}

	return response.Data, nil
}
