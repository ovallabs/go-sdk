package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const customerAPIVersion = "v1/customer"

// CreateCustomer makes an API request using Call to create a customer
func (c *Call) CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (model.Customer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, customerAPIVersion)

	fL := c.logger.With().Str("func", "CreateCustomer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("email", request.Email).Str("name", request.Name).
		Str("mobileNumber", request.MobileNumber).Str("reference", request.Reference).
		Str("yieldOfferingId", request.YieldOfferingID).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Customer `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Customer{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.Customer{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// UpdateCustomer makes an API request using Call to update a customer
func (c *Call) UpdateCustomer(ctx context.Context, request model.UpdateCustomerRequest) (model.Customer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, customerAPIVersion)

	fL := c.logger.With().Str("func", "UpdateCustomer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("email", request.Email).Str("name", request.Name).
		Str("mobileNumber", request.MobileNumber).Str("reference", request.Reference).
		Str("yieldOfferingId", request.MobileNumber).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.Customer `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Patch(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Customer{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.Customer{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetAllCustomers makes an API request using Call to get all customers
func (c *Call) GetAllCustomers(ctx context.Context) ([]model.Customer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, customerAPIVersion)

	fL := c.logger.With().Str("func", "GetAllCustomers").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []model.Customer `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return []model.Customer{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return []model.Customer{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetCustomerByID makes an API request using Call to get a customer by ID
func (c *Call) GetCustomerByID(ctx context.Context, request model.GetCustomerByIDRequest) (model.CustomerInfo, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, customerAPIVersion, request.CustomerID)

	fL := c.logger.With().Str("func", "GetCustomerByID").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.CustomerInfo `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.CustomerInfo{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.CustomerInfo{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
