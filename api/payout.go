package api

import (
	"context"
	"fmt"
	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
	"github.com/rs/zerolog/log"
	"net/http"
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
