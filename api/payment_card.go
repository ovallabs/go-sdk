package api

import (
	"context"
	"fmt"

	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

// InitiatePaymentCardRequest makes request to Torus to initiate a payment card for a customer
func (c *Call) InitiatePaymentCardRequest(ctx context.Context, request model.InitiateCardRequest) (string, error) {
	var (
		err       error
		response  string
		path      = "v1/payments/cards/initiate"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}

// CompletePaymentCardRequest makes request to Torus to complete a payment card for a customer
func (c *Call) CompletePaymentCardRequest(ctx context.Context, request model.CompleteCardRequest) error {
	var (
		err  error
		path = "v1/payments/cards/complete"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, nil)

	return err
}

// GetLinkToAddPaymentCard makes request to Torus to add a payment card
func (c *Call) GetLinkToAddPaymentCard(ctx context.Context, request model.GetLinkToAddCardReq) (string, error) {
	var (
		err      error
		response string
		path     = "v1/payments/cards"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetLinkToAuthorizeCustomer makes request to Torus to get link to authorize customer for payment card
func (c *Call) GetLinkToAuthorizeCustomer(ctx context.Context, request model.GetLinkToAddCardReq) (string, error) {
	var (
		err      error
		response string
		path     = "v1/payments/cards/authorize"
	)

	err = c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, &response)

	return response, err
}

// GetCustomerPaymentCards makes request to Torus to get al payment cards for a customer
func (c *Call) GetCustomerPaymentCards(ctx context.Context, customerID string, status, search *string, dateBetween *model.DateBetween, page *model.Page) (model.AllPaymentCardsResponse, error) {
	var (
		err      error
		response model.AllPaymentCardsResponse
		params   = make(map[string]interface{})
		path     = fmt.Sprintf("v1/payments/cards/%s", customerID)
	)

	if status != nil {
		params["status"] = *status
	}
	if search != nil {
		params["search"] = *search
	}
	if dateBetween != nil {
		helpers.FillParamsWithDateInterval(params, *dateBetween)
	}
	if page != nil {
		helpers.FillParamsWithPage(params, *page)
	}

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)

	return response, err
}

// GetCustomerPaymentCardByID makes request to Torus to get link to authorize customer for payment card by the ID
func (c *Call) GetCustomerPaymentCardByID(ctx context.Context, customerID, ID string) (model.PaymentCard, error) {
	var (
		err      error
		response model.PaymentCard
		path     = fmt.Sprintf("v1/payments/cards/%s/%s", customerID, ID)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, nil, nil, nil, &response)

	return response, err
}

// DebitPaymentCard makes request to Torus to debit a customer payment card
func (c *Call) DebitPaymentCard(ctx context.Context, request model.DebitCustomerPaymentCardRequest) (string, error) {
	var (
		err       error
		response  string
		path      = "v1/payments/cards/debit"
		signature = helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	)

	err = c.makeRequest(ctx, path, http.MethodPost, &signature, nil, nil, request, &response)

	return response, err
}
