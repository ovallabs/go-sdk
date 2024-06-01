package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const currencySwapAPIVersion = "v1/currency-swaps"

// InitiateCurrencySwap makes a request to Torus to initiate currency swap
func (c *Call) InitiateCurrencySwap(ctx context.Context, request model.InitiateCurrencySwapRequest) (model.CurrencySwap, error) {
	var (
		err      error
		response model.CurrencySwap
		path     = currencySwapAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetCurrencySwaps makes a request to Torus to get all currency swaps
func (c *Call) GetCurrencySwaps(ctx context.Context, status, from, to string, dateBetween model.DateBetween, page model.Page) (model.AllSwapsResponse, error) {
	var (
		err      error
		response model.AllSwapsResponse
		params   = make(map[string]interface{})
		path     = currencySwapAPIVersion
	)

	if status != "" {
		params["status"] = status
	}
	if to != "" {
		params["to_currency"] = to
	}
	if from != "" {
		params["from_currency"] = from
	}
	if dateBetween != (model.DateBetween{}) {
		helpers.FillParamsWithDateInterval(params, dateBetween)
	}
	if page != (model.Page{}) {
		helpers.FillParamsWithPage(params, page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetCurrencySwapByID makes a request to Torus to get currency swap by ID
func (c *Call) GetCurrencySwapByID(ctx context.Context, currencySwapID string) (model.CurrencySwap, error) {
	var (
		err      error
		response model.CurrencySwap
		path     = fmt.Sprintf("%s/%s", currencySwapAPIVersion, currencySwapID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
