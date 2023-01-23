package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const customerAPIVersion = "v1/customer"

// CreateCustomer makes an API request using Call to create a customer
func (c *Call) CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (model.Customer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, customerAPIVersion)

	fL := c.logger.With().Str("func", "CreateCustomer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
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
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
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
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
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
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
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
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
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
	fL.Info().Str("customerID", request.CustomerID).Interface(model.LogStrRequest, "empty").Msg("request")
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
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.CustomerInfo{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetCustomerBalance to get customer balances
func (c Call) GetCustomerBalance(ctx context.Context, request model.GetCustomerBalanceRequest) (model.CustomerBalanceResponse, error) {
	endpoint := fmt.Sprintf("%s%s%s?customer_id=%s&yield_offering_id=%s", c.baseURL, customerAPIVersion, "/balance", request.CustomerID, request.YieldOfferingID)

	fL := c.logger.With().Str("func", "GetCustomerBalance").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("customerID", request.CustomerID.String()).Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.CustomerBalanceResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.CustomerBalanceResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.CustomerBalanceResponse{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetCustomerBalances to get customer balances from different yield offering
func (c Call) GetCustomerBalances(ctx context.Context, customerID uuid.UUID) (model.CustomerBalancesResponse, error) {
	endpoint := fmt.Sprintf("%s%s%s/%s", c.baseURL, customerAPIVersion, "/balances", customerID)

	fL := c.logger.With().Str("func", "GetCustomerBalances").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("customerID", customerID.String()).Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.CustomerBalancesResponse `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.CustomerBalancesResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.CustomerBalancesResponse{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

func (c Call) DeleteCustomer(ctx context.Context, customerID uuid.UUID) error {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, customerAPIVersion, customerID)

	fL := c.logger.With().Str("func", "DeleteCustomer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("customerID", customerID.String()).Interface(model.LogStrRequest, "empty").Msg("request")

	defer fL.Info().Msg("done...")

	var response interface{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Delete(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response).Msg("response")

	return nil
}
