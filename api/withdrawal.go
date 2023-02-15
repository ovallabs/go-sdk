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
	// extract request id value from context
	requestID := helpers.GetRequestID(ctx)

	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeaders(map[string]string{
			"Signature":              signature,
			model.RequestIDHeaderKey: requestID,
		}).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Withdrawal{}, model.ErrNetworkError
		}
		return model.Withdrawal{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// FiatWithdrawal makes an API request to withdrawal to a provided bank account
func (c *Call) FiatWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, withdrawalAPIVersion, "/fiat")

	fL := c.logger.With().Str("func", "FiatWithdrawal").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	// extract request id value from context
	requestID := helpers.GetRequestID(ctx)

	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeaders(map[string]string{
			"Signature":              signature,
			model.RequestIDHeaderKey: requestID,
		}).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Withdrawal{}, model.ErrNetworkError
		}
		return model.Withdrawal{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// CryptoWithdrawal makes an API request to withdrawal to a specified crypto wallet address
func (c Call) CryptoWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, withdrawalAPIVersion, "/crypto")

	fL := c.logger.With().Str("func", "CryptoWithdrawal").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	// extract request id value from context
	requestID := helpers.GetRequestID(ctx)

	response := struct {
		Data model.Withdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeaders(map[string]string{
			"Signature":              signature,
			model.RequestIDHeaderKey: requestID,
		}).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Withdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Withdrawal{}, model.ErrNetworkError
		}
		return model.Withdrawal{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// FeeWithdrawal makes an API request to withdrawal for withdrawal fees
func (c *Call) FeeWithdrawal(ctx context.Context, request model.FeeWithdrawalRequest) (model.FeeWithdrawal, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, withdrawalAPIVersion, "/fee")

	fL := c.logger.With().Str("func", "FeeWithdrawal").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	// extract request id value from context
	requestID := helpers.GetRequestID(ctx)

	response := struct {
		Data model.FeeWithdrawal `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeaders(map[string]string{
			"Signature":              signature,
			model.RequestIDHeaderKey: requestID,
		}).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.FeeWithdrawal{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.FeeWithdrawal{}, model.ErrNetworkError
		}
		return model.FeeWithdrawal{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}
