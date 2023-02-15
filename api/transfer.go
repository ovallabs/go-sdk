package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const transferAPIVersion = "v1/transfer"

// InitiateTransfer makes an API request using Call to initiate a transfer
func (c *Call) InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.Transfer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, transferAPIVersion)

	fL := c.logger.With().Str("func", "InitiateTransfer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)

	response := struct {
		Data model.Transfer `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeaders(map[string]string{
			"Signature":              signature,
			model.RequestIDHeaderKey: helpers.GetRequestID(ctx),
		}).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Transfer{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Transfer{}, model.ErrNetworkError
		}
		return model.Transfer{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetExchangeRates makes an API request using Call to get exchange rates
func (c *Call) GetExchangeRates(ctx context.Context, request model.GetExchangeRateRequest) (model.ExchangeRateDetails, error) {
	endpoint := fmt.Sprintf("%s%s/quote?amount=%f&source_currency=%s&destination_currency=%s",
		c.baseURL, transferAPIVersion, request.Amount, request.SourceCurrency, request.DestinationCurrency)

	fL := c.logger.With().Str("func", "GetExchangeRates").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.ExchangeRateDetails `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.ExchangeRateDetails{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.ExchangeRateDetails{}, model.ErrNetworkError
		}
		return model.ExchangeRateDetails{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
