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
	depositAPIVersion       = "v1/deposit"
	depositsAPIVersion      = "v1/deposits"
	fundTransferAPIVersion  = "v1/transfer-funds"
	intraTransferAPIVersion = "v1/intra-transfer"
)

// InitiateDeposit makes an API request using Call to initiate a deposit
func (c *Call) InitiateDeposit(ctx context.Context, request model.InitiateDepositRequest) (model.Deposit, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, depositAPIVersion)

	fL := c.logger.With().Str("func", "InitiateDeposit").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Deposit `json:"data"`
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
		return model.Deposit{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Deposit{}, model.ErrNetworkError
		}
		return model.Deposit{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetAllDeposits makes an API request using Call to get all deposits
func (c *Call) GetAllDeposits(ctx context.Context) (model.DepositBatchResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, depositsAPIVersion)

	fL := c.logger.With().Str("func", "GetAllDeposits").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.DepositBatchResponse `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.DepositBatchResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.DepositBatchResponse{}, model.ErrNetworkError
		}
		return model.DepositBatchResponse{}, model.ParseError(errRes.Error.Details)
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetDepositID makes an API request using Call to get deposit by id
func (c *Call) GetDepositID(ctx context.Context, id uuid.UUID) (model.Deposit, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, depositAPIVersion, id)

	fL := c.logger.With().Str("func", "GetDepositID").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.Deposit `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Deposit{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Deposit{}, model.ErrNetworkError
		}
		return model.Deposit{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// InternalFundsTransfer makes an api
func (c *Call) InternalFundsTransfer(ctx context.Context, request model.FundTransferRequest) (model.Deposit, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, fundTransferAPIVersion)

	fL := c.logger.With().Str("func", "InternalFundsTransfer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Deposit `json:"data"`
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
		return model.Deposit{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.Deposit{}, model.ErrNetworkError
		}
		return model.Deposit{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}

// IntraTransfer makes an API request to transfer funds between two customers
func (c *Call) IntraTransfer(ctx context.Context, request model.IntraTransferRequest) (model.IntraTransferResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, intraTransferAPIVersion)

	fL := c.logger.With().Str("func", "IntraTransfer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface("request", request).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.IntraTransferResponse `json:"data"`
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
		return model.IntraTransferResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		var errRes model.ErrorResponse
		errRes, err = model.GetErrorDetails(string(res.Body()))
		if err != nil {
			fL.Err(err).Msg("error occurred")
			return model.IntraTransferResponse{}, model.ErrNetworkError
		}
		return model.IntraTransferResponse{}, model.ParseError(errRes.Error.Details)
	}

	return response.Data, nil
}
