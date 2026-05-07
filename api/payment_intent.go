package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

// CreateCustomerPaymentIntent makes request to create a customer payment intent.
func (c *Call) CreateCustomerPaymentIntent(ctx context.Context, request model.CreateCustomerPaymentIntentRequest) (model.CreateCustomerPaymentIntentResponse, error) {
	var (
		err      error
		response model.CreateCustomerPaymentIntentResponse
		path     = "v1/payments/intents"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// CompleteCustomerPaymentIntent makes request to complete a customer payment intent.
func (c *Call) CompleteCustomerPaymentIntent(ctx context.Context, request model.CompleteCustomerPaymentIntentRequest) (model.CreateCustomerPaymentIntentResponse, error) {
	var (
		err      error
		response model.CreateCustomerPaymentIntentResponse
		path     = "v1/payments/intents/complete"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// AuthenticateCustomerPaymentIntent makes request to authenticate a customer payment intent.
func (c *Call) AuthenticateCustomerPaymentIntent(ctx context.Context, request model.AuthenticateCustomerPaymentIntentRequest) (model.CreateCustomerPaymentIntentResponse, error) {
	var (
		err      error
		response model.CreateCustomerPaymentIntentResponse
		path     = "v1/payments/intents/authenticate"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetCustomerPaymentIntentByID makes request to get a customer payment intent by ID.
func (c *Call) GetCustomerPaymentIntentByID(ctx context.Context, paymentIntentID string) (model.CustomerPaymentIntent, error) {
	var (
		err      error
		response model.CustomerPaymentIntent
		path     = fmt.Sprintf("v1/payments/intents/%s", paymentIntentID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}
