package api

import (
	"context"
	"fmt"
	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
	"net/http"
)

const transferAPIVersion = "v1/transfer"

// InitiateTransfer makes an API request using Call to initiate a transfer
func (c *Call) InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.Transfer, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, transferAPIVersion)

	fL := c.logger.With().Str("func", "InitiateTransfer").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("customerID", request.CustomerID).Str("reference", request.Reference).
		Str("currency", request.Currency).Interface("destination", request.Destination).
		Float64("amount", request.Amount).Str("note", request.Note).Str("reason", request.Reason).
		Str("reference", request.Reference).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.Transfer `json:"data"`
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
		return model.Transfer{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.Transfer{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
