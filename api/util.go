package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/go-resty/resty/v2"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

func (c *Call) makeRequest(ctx context.Context, path, method string, params, formData map[string]interface{}, requestBody, responseData interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")

	var (
		err             error
		res             *resty.Response
		genericResponse model.GenericResponse
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
		err = json.Unmarshal([]byte(result), &genericResponse)
		if err != nil {
			log.Err(err).Msg("error decoding response")
			return err
		}
		if data, ok := genericResponse.Data.(map[string]interface{}); !ok {
			err := func(v1, v2 interface{}) error {
				value := reflect.ValueOf(v1)
				if value.Kind() != reflect.Ptr {
					return errors.New("responseData must be a pointer")
				}
				value = value.Elem()
				if value.Kind() != reflect.TypeOf(v2).Kind() {
					return errors.New("responseData does not match returned data")
				}
				value.Set(reflect.ValueOf(v2))
				return nil
			}(responseData, genericResponse.Data)
			if err != nil {
				log.Err(err).Msg(err.Error())
				return err
			}
		} else {
			jsonData, _ := json.Marshal(data)
			_ = json.Unmarshal(jsonData, &responseData)
		}
	} else if res.StatusCode() >= 400 {
		result := string(res.Body())
		err = json.Unmarshal([]byte(result), &genericResponse)
		if err != nil {
			log.Err(err).Msg("error decoding response")
			return err
		}
		err = errors.New(genericResponse.Error.Details)
		log.Err(err).Msg("error occurred while processing request")
		return err
	}

	log.Info().Interface(model.LogStrResponse, responseData).Msg("response")
	return nil
}
