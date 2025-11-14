package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

const cryptoAPIVersion = "v1/crypto"

// GetCustomerWallet makes request to Torus to get customer wallet if available or  create customer wallet if not
func (c *Call) GetCustomerWallet(ctx context.Context, request model.CustomerWalletRequest) (model.CustomerWallet, error) {
	var (
		err      error
		response model.CustomerWallet
		params   = map[string]interface{}{"customer_id": request.CustomerID, "network": request.Network, "asset": request.Asset}
		path     = fmt.Sprintf("%s/wallet", cryptoAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetSupportedAssets makes request to Torus to get supported assets
func (c *Call) GetSupportedAssets(ctx context.Context) ([]*model.SupportedCurrencies, error) {
	var (
		err      error
		response []*model.SupportedCurrencies
		path     = "v1/supported-assets"
	)
	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)
	return response, err
}
