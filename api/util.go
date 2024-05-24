package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

func (c *Call) makeRequest(ctx context.Context, path, method string, params, formData map[string]interface{}, requestBody, responseBody interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")

	var (
		err           error
		res           *resty.Response
		errorResponse model.ErrorResponse
	)

	client := c.client.R().
		SetAuthToken(c.bearerToken).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetContext(ctx)

	if requestBody != nil {
		log.Info().Interface(model.LogStrRequest, requestBody).Msg("request")
		client.SetBody(requestBody)
	}

	if params != nil {
		log.Info().Interface(model.LogStrParams, params).Msg("parameters")
		for k, v := range params {
			client.SetQueryParam(k, v.(string))
		}
	}

	if formData != nil {
		log.Info().Interface(model.LogStrForm, formData).Msg("form data")
		formDataConv := make(map[string]string)
		for k, v := range formData {
			if file, ok := v.(*os.File); ok {
				client.SetFileReader(k, file.Name(), file)
			} else {
				formDataConv[k] = v.(string)
			}
		}
		client.SetFormData(formDataConv)
	}

	switch method {
	case http.MethodGet:
		res, err = client.Get(endpoint)
	case http.MethodPost:
		res, err = client.Post(endpoint)
	case http.MethodPut:
		res, err = client.Put(endpoint)
	case http.MethodPatch:
		res, err = client.Patch(endpoint)
	case http.MethodHead:
		res, err = client.Head(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return err
	}

	if err != nil {
		log.Err(err).Msg("something went wrong")
		if res != nil {
			log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		}
		return err
	}

	if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		result := string(res.Body())
		err = json.Unmarshal([]byte(result), responseBody)
		if err != nil {
			log.Err(err).Msg("error decoding response")
		}
	} else if res.StatusCode() >= 400 {
		result := string(res.Body())
		err = json.Unmarshal([]byte(result), &errorResponse)
		if err != nil {
			log.Err(err).Msg("error decoding response")
			return err
		}
		err = errors.New(errorResponse.Message)
		log.Err(err).Msg("error occurred while processing request")
		return err
	}

	log.Info().Interface(model.LogStrResponse, responseBody).Msg("response")
	return nil
}
