package api

import (
	"context"
	"fmt"
	"github.com/ovalfi/go-sdk/model"
	"net/http"
)

const depositAccountAPIVersion = "v1/deposit-accounts"

// GetDepositAccount makes request to Torus to get deposit account
func (c *Call) GetDepositAccount(ctx context.Context, currency string) (model.AccountDetails, error) {
	var (
		err      error
		response model.AccountDetails
		path     = fmt.Sprintf("%s/%s", depositAccountAPIVersion, currency)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
