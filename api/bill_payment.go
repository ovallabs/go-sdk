package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const billPaymentAPIVersion = "v1/bills"

// GetBillerCategories makes a request to Torus to get the list of bill payment categories
func (c *Call) GetBillerCategories(ctx context.Context) ([]model.BillerCategory, error) {
	var (
		err      error
		response []model.BillerCategory
		path     = fmt.Sprintf("%s/categories", billPaymentAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetBillers makes a request to Torus to get the list of billers configured under a category
func (c *Call) GetBillers(ctx context.Context, category string) ([]model.Biller, error) {
	var (
		err      error
		response []model.Biller
		path     = fmt.Sprintf("%s/categories/%s/billers", billPaymentAPIVersion, category)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetBillerProducts makes a request to Torus to get the list of products offered by a biller
func (c *Call) GetBillerProducts(ctx context.Context, category, biller string, paymentType *string, page *model.Page) (model.AllBillerProductsResponse, error) {
	var (
		err      error
		response model.AllBillerProductsResponse
		params   = make(map[string]interface{})
		path     = fmt.Sprintf("%s/categories/%s/billers/%s/products", billPaymentAPIVersion, category, biller)
	)

	if paymentType != nil {
		params["payment_type"] = *paymentType
	}
	if page != nil {
		helpers.FillParamsWithPage(params, *page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// ValidateBillerCustomer makes a request to Torus to validate a customer's identifier
// (e.g. meter or smart card number) against a biller product before payment.
func (c *Call) ValidateBillerCustomer(ctx context.Context, request model.ValidateBillerCustomerRequest) (model.ValidateBillerCustomerResponse, error) {
	var (
		err      error
		response model.ValidateBillerCustomerResponse
		path     = fmt.Sprintf("%s/validate-customer", billPaymentAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// PayBill makes a request to Torus to initiate a bill payment
func (c *Call) PayBill(ctx context.Context, request model.PayBillRequest) (model.PayBillResponse, error) {
	var (
		err      error
		response model.PayBillResponse
		path     = fmt.Sprintf("%s/pay", billPaymentAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetBillPaymentTransaction makes a request to Torus to get a bill payment transaction by its ID
func (c *Call) GetBillPaymentTransaction(ctx context.Context, billPaymentID string) (model.BillPaymentTransaction, error) {
	var (
		err      error
		response model.BillPaymentTransaction
		path     = fmt.Sprintf("%s/payments/%s", billPaymentAPIVersion, billPaymentID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
