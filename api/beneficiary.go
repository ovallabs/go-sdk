package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const beneficiaryAPIVersion = "v1/beneficiaries"

// CreateBeneficiary makes request to Torus to create beneficiary
func (c *Call) CreateBeneficiary(ctx context.Context, request model.CreateBeneficiaryRequest) (model.TransferBeneficiary, error) {
	var (
		err      error
		response model.TransferBeneficiary
		path     = beneficiaryAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetBeneficiaries makes request to Torus to get all beneficiaries
func (c *Call) GetBeneficiaries(ctx context.Context, currency string, page *model.Page) (model.AllBeneficiariesResponse, error) {
	var (
		err      error
		response model.AllBeneficiariesResponse
		params   = make(map[string]interface{})
		path     = beneficiaryAPIVersion
	)

	if currency != "" {
		params["destination_currency"] = currency
	}
	if page != nil {
		helpers.FillParamsWithPage(params, *page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetBeneficiaryByID makes request to Torus to get beneficiary by its ID
func (c *Call) GetBeneficiaryByID(ctx context.Context, beneficiaryID string) (model.TransferBeneficiary, error) {
	var (
		err      error
		response model.TransferBeneficiary
		path     = fmt.Sprintf("%s/%s", beneficiaryAPIVersion, beneficiaryID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
