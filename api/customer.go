package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const customerAPIVersion = "v1/customer"

// CreateCustomer makes request to Torus to create customer
func (c *Call) CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (model.Customer, error) {
	var (
		err       error
		response  model.Customer
		path      = customerAPIVersion
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// UpdateCustomer makes request to Torus to update customer
func (c *Call) UpdateCustomer(ctx context.Context, request model.UpdateCustomerRequest) (model.Customer, error) {
	var (
		err      error
		response model.Customer
		path     = customerAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodPatch, nil, nil, nil, request, &response)

	return response, err
}

// GetAllCustomers makes request to Torus to get all customers
func (c *Call) GetAllCustomers(ctx context.Context) (model.AllCustomersResponse, error) {
	var (
		err      error
		response model.AllCustomersResponse
		path     = customerAPIVersion
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetCustomerByID makes request to Torus to get customer by its ID
func (c *Call) GetCustomerByID(ctx context.Context, customerID string) (model.Customer, error) {
	var (
		err      error
		response model.Customer
		path     = fmt.Sprintf("%s/%s", customerAPIVersion, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// GetCustomerBalance makes request to Torus to get customer balance from a yield offering
func (c *Call) GetCustomerBalance(ctx context.Context, customerID, yieldOfferingID string) (model.CustomerBalance, error) {
	var (
		err      error
		response model.CustomerBalance
		params   = map[string]interface{}{"customer_id": customerID, "yield_offering_id": yieldOfferingID}
		path     = fmt.Sprintf("%s/balance", customerAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetCustomerBalances makes request to Torus to get customer balance from different yield offering
func (c *Call) GetCustomerBalances(ctx context.Context, customerID string) (model.CustomerBalances, error) {
	var (
		err      error
		response model.CustomerBalances
		path     = fmt.Sprintf("%s/balances/%s", customerAPIVersion, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// DeleteCustomer makes request to Torus to delete customer
func (c *Call) DeleteCustomer(ctx context.Context, customerID string) error {
	var (
		err  error
		path = fmt.Sprintf("%s/%s", customerAPIVersion, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, nil, nil, nil, nil)

	return err
}
