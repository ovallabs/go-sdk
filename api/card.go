package api

import (
	"context"
	"fmt"

	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

// CreateCustomerCard makes request to Torus to create a card for a customer
func (c *Call) CreateCustomerCard(ctx context.Context, request model.CreateCustomerCardRequest) (string, error) {
	var (
		err       error
		response  string
		path      = "v1/cards"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)
	return response, err
}

// FreezeUnfreezeCard makes request to Torus to freeze/unfreeze a customer card
func (c *Call) FreezeUnfreezeCard(ctx context.Context, request model.FreezeCardRequest) (string, error) {
	var (
		err      error
		response string
		path     = "v1/cards/freeze"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)
	return response, err
}

// GetCustomerCards makes request to Torus to get customer cards
func (c *Call) GetCustomerCards(ctx context.Context, customerID *string) (model.AllCardsResponse, error) {
	var (
		err      error
		response model.AllCardsResponse
		path     = "v1/cards"
	)

	if customerID != nil {
		path = fmt.Sprintf("v1/cards?customer_id=%s", *customerID)
	}
	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)
	return response, err
}

// GetCustomerCardByID makes request to Torus to get customer card by ID
func (c *Call) GetCustomerCardByID(ctx context.Context, cardID string) (model.Card, error) {
	var (
		err      error
		response model.Card
		path     = fmt.Sprintf("v1/cards/%s", cardID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)
	return response, err
}

// FundCustomerCard makes request to Torus to fund a card for a customer
func (c *Call) FundCustomerCard(ctx context.Context, request model.FundCustomerCardRequest) (model.Card, error) {
	var (
		err      error
		response model.Card
		path     = "v1/cards/fund"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)
	return response, err
}

// GetCustomerCardSecureDetails makes request to Torus to get customer card secure details
func (c *Call) GetCustomerCardSecureDetails(ctx context.Context, cardID, customerID string) (model.VaultedCardDetails, error) {
	var (
		err      error
		response model.VaultedCardDetails
		path     = fmt.Sprintf("v1/cards/%s/secure?customer_id=%s", cardID, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)
	return response, err
}

// DeleteCard makes request to Torus to delete/terminate a customer card
func (c *Call) DeleteCard(ctx context.Context, cardID, customerID string) (string, error) {
	var (
		err      error
		response string
		path     = fmt.Sprintf("v1/cards/%s?customer_id=%s", cardID, customerID)
	)

	err = c.makeRequest(ctx, path, http.MethodDelete, nil, nil, nil, nil, &response)
	return response, err
}

// InitiateCustomerPaymentSession makes request to torus to initiate a customer payment session
func (c *Call) InitiateCustomerPaymentSession(ctx context.Context, request model.CustomerPaymentSessionRequest) (model.CustomerPaymentSessionResponse, error) {
	var (
		err      error
		response model.CustomerPaymentSessionResponse
		path     = fmt.Sprintf("v1/cards/initiate-customer-payment-session")
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, response)

	return response, err
}
