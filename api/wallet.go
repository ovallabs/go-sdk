package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const (
	walletAPIVersion = "v1/payments/crypto/wallet"
	assetAPIVersion  = "v1/payments/crypto/supported-assets"
)

// GetWallet makes an API request using Call to get customer wallet
func (c *Call) GetWallet(ctx context.Context, request model.WalletRequest) (model.Wallet, error) {
	endpoint := fmt.Sprintf("%s%s?customer_id=%s&network=%s&asset=%s", c.baseURL, walletAPIVersion, request.CustomerID, request.Network, request.Asset)

	fL := c.logger.With().Str("func", "GetWallet").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.Wallet `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Wallet{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Wallet{}, model.ErrNetworkError
		}
		return model.Wallet{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetWallets makes an API request using Call to get all customers wallets
func (c *Call) GetWallets(ctx context.Context, customerID string) ([]*model.Wallet, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, walletAPIVersion, customerID)

	fL := c.logger.With().Str("func", "GetWallets").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, customerID).Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []*model.Wallet `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return nil, model.ErrNetworkError
		}
		return nil, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetSupportedAssets makes an API request using Call to get all supported assets
func (c *Call) GetSupportedAssets(ctx context.Context) ([]*model.SupportedAsset, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, assetAPIVersion)

	fL := c.logger.With().Str("func", "GetSupportedAssets").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []*model.SupportedAsset `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return nil, model.ErrNetworkError
		}
		return nil, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
