package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

// GetPayoutByID makes a request to Torus to get the payout by its ID.
func (c *Call) GetPayoutByID(ctx context.Context, payoutID string) (model.PayoutResponse, error) {
	endpoint := fmt.Sprintf("%s/payouts/%s", c.baseURL, payoutID)

	fL := c.logger.With().Str("func", "GetPayoutByID").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, payoutID).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data model.PayoutResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetResult(&response).
		SetError(&errorRes).
		Get(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return model.PayoutResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return model.PayoutResponse{}, err
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")

	return response.Data, err
}

// InitiateDirectBulkPayout makes a request to Torus to initiate a bulk payout
func (c *Call) InitiateDirectBulkPayout(ctx context.Context, request model.InitiatePayoutRequest) (model.PayoutDetails, error) {
	endpoint := fmt.Sprintf("%s/payouts", c.baseURL)

	fL := c.logger.With().Str("func", "InitiateDirectBulkPayout").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data model.PayoutDetails `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetBody(request).
		SetResult(&response).
		SetError(&errorRes).
		Post(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return model.PayoutDetails{}, err
	}

	if res.StatusCode() != http.StatusCreated {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return model.PayoutDetails{}, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, err
}

// InitiatePayout makes a request to Torus to initiate a bulk payout
func (c *Call) InitiatePayout(ctx context.Context, currency, payoutType, beneficiaryType, remarks, filePath string) (model.PayoutDetails, error) {
	endpoint := fmt.Sprintf("%s/payouts/upload", c.baseURL)

	formData := map[string]string{
		"currency":         currency,
		"payout_type":      payoutType,
		"beneficiary_type": beneficiaryType,
		"remarks":          remarks,
	}

	fL := c.logger.With().Str("func", "InitiatePayout").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, formData).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data model.PayoutDetails `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetFile("document", filePath).
		SetFormData(formData).
		SetResult(&response).
		SetError(&errorRes).
		Post(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return model.PayoutDetails{}, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return model.PayoutDetails{}, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, nil
}

// GetAllPayouts makes request to Torus to get all payouts
func (c *Call) GetAllPayouts(ctx context.Context, status, search string, dateBetween model.DateBetween, page model.Page) (model.AllPayoutsResponse, error) {
	endpoint := fmt.Sprintf("%s/payouts", c.baseURL)

	params := map[string]string{}
	if status != "" {
		params["status"] = status
	}
	if search != "" {
		params["search"] = search
	}
	if dateBetween != (model.DateBetween{}) {
		if dateBetween.From != "" {
			params["from"] = dateBetween.From
		}
		if dateBetween.To != "" {
			params["to"] = dateBetween.To
		}
	}
	if page != (model.Page{}) {
		if page.Number != nil {
			params["number"] = strconv.Itoa(*page.Number)
		}
		if page.Size != nil {
			params["size"] = strconv.Itoa(*page.Size)
		}
		if page.SortBy != "" {
			params["sort_by"] = page.SortBy
		}
		if page.SortDirectionDesc != nil {
			params["sort_direction_desc"] = strconv.FormatBool(*page.SortDirectionDesc)
		}
	}

	fL := c.logger.With().Str("func", "GetAllPayouts").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, params).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data model.AllPayoutsResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetQueryParams(params).
		SetResult(&response).
		SetError(&errorRes).
		Get(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return model.AllPayoutsResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return model.AllPayoutsResponse{}, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, nil
}

// CancelPayout makes request to Torus to cancel payout
func (c *Call) CancelPayout(ctx context.Context, request model.CancelPayoutRequest) (bool, error) {
	endpoint := fmt.Sprintf("%s/payouts/cancel", c.baseURL)

	fL := c.logger.With().Str("func", "CancelPayout").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data bool `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetBody(request).
		SetResult(&response).
		SetError(&errorRes).
		Post(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return false, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return false, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, nil
}

// UpdatePayoutAccount makes request to Torus to update payout account by its ID
func (c *Call) UpdatePayoutAccount(ctx context.Context, payoutID string, request model.TransferBeneficiaryDetails) (bool, error) {
	endpoint := fmt.Sprintf("%s/payouts/accounts/%s", c.baseURL, payoutID)

	fL := c.logger.With().Str("func", "UpdatePayoutAccount").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data bool `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetBody(request).
		SetResult(&response).
		SetError(&errorRes).
		Put(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return false, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return false, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, nil
}

// GetPayoutConfig makes request to Torus to get payout config
func (c *Call) GetPayoutConfig(ctx context.Context, currency string) (model.BulkPayoutConfig, error) {
	endpoint := fmt.Sprintf("%s/payouts/config/%s", c.baseURL, currency)

	fL := c.logger.With().Str("func", "GetPayoutConfig").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, currency).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data model.BulkPayoutConfig `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetResult(&response).
		SetError(&errorRes).
		Get(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return model.BulkPayoutConfig{}, err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return model.BulkPayoutConfig{}, model.ErrNetworkError
	}

	log.Info().Interface(model.LogStrResponse, response).Msg("response returned")
	return response.Data, nil
}

// GetPayoutDocumentTemplate makes request to Torus to get payout document template
func (c *Call) GetPayoutDocumentTemplate(ctx context.Context, currency, docType string) (string, error) {
	endpoint := fmt.Sprintf("%s/payouts/template", c.baseURL)

	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	if docType != "" {
		params["type"] = docType
	}

	fL := c.logger.With().Str("func", "GetPayoutDocumentTemplate").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, params).Msg("request")
	defer fL.Info().Msg("done...")

	errorRes := model.ErrorResponse{}
	response := struct {
		Data string `json:"data"`
	}{}

	res, err := c.client.R().
		SetHeader("request_id", helpers.GetRequestID(ctx)).
		SetAuthToken(c.bearerToken).
		SetContext(ctx).
		SetQueryParams(params).
		SetResult(&response).
		SetError(&errorRes).
		Get(endpoint)

	if err != nil {
		log.Err(err).Msg("something went wrong with this request")
		return "", err
	}

	if res.StatusCode() != http.StatusOK {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msgf(" %+v", errorRes)
		return "", model.ErrNetworkError
	}

	return response.Data, nil
}
