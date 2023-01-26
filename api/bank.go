package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const (
	bankEndpoint        = "v1/payments/banks"
	bankAccountEndpoint = "v1/payments/banks/account"
)

// ResolveBankAccount makes an API request using Call to resolve bank account
func (c Call) ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetailResponse, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, bankEndpoint, "/resolve-account")

	fL := c.logger.With().Str("func", "ResolveBankAccount").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.AccountDetailResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.AccountDetailResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.AccountDetailResponse{}, model.ErrNetworkError
		}
		return model.AccountDetailResponse{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// GetBanks makes an API request using Call to get list of banks
func (c Call) GetBanks(ctx context.Context) ([]model.BankCodeResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, bankEndpoint)

	fL := c.logger.With().Str("func", "GetBanks").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []model.BankCodeResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return response.Data, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return []model.BankCodeResponse{}, model.ErrNetworkError
		}
		return []model.BankCodeResponse{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// GenerateBankAccount makes an API request to generate bank account
func (c *Call) GenerateBankAccount(ctx context.Context, request model.BankAccountRequest) (model.BankAccountResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, bankAccountEndpoint)

	fL := c.logger.With().Str("func", "GenerateBankAccount").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.BankAccountResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		return model.BankAccountResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.BankAccountResponse{}, model.ErrNetworkError
		}
		return model.BankAccountResponse{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// GetBankAccount makes an API request to get bank account
func (c *Call) GetBankAccount(ctx context.Context, customerID uuid.UUID) (model.BankAccountResponse, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, bankAccountEndpoint, customerID)

	fL := c.logger.With().Str("func", "GetBankAccount").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", customerID.String()).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.BankAccountResponse `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.BankAccountResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.BankAccountResponse{}, model.ErrNetworkError
		}
		return model.BankAccountResponse{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}
