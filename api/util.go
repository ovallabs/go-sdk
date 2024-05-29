package api

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

func (c *Call) makeRequest(ctx context.Context, path, method string, signature *string, params, formData map[string]interface{}, requestBody, responseData interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")

	var (
		err             error
		res             *resty.Response
		genericResponse = model.GenericResponse{}
	)

	client := c.client.R().
		SetAuthToken(c.bearerToken).
		SetHeader(model.RequestIDHeaderKey, helpers.GetRequestID(ctx)).
		SetResult(&genericResponse).
		SetError(&genericResponse).
		SetContext(ctx)

	if signature != nil {
		client = client.SetHeader("Signature", *signature)
	}

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
				name := file.Name()
				contentType := mime.TypeByExtension(filepath.Ext(name))
				client.SetMultipartField(k, name, contentType, file)
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

	if genericResponse.Error != nil {
		err = errors.New(genericResponse.Error.Details)
		log.Err(err).Msg("error while making request")
		return err
	}

	log.Info().Interface(model.LogStrResponse, genericResponse.Data).Msg("response")
	if responseData != nil {
		err = mapstruct(genericResponse.Data, responseData)
		if err != nil {
			return err
		}
	}
	return nil
}

// mapstruct map api call result to the expected interface
func mapstruct(data, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			stringToTimeHookFunc(),
			stringToUUIDHookFunc(),
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(data)

	return err
}

// stringToUUIDHookFunc type conversion for string to uuid.UUID
func stringToUUIDHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() == reflect.String && t == reflect.TypeOf(uuid.UUID{}) {
			str, ok := data.(string)
			if !ok {
				return data, nil
			}
			return uuid.Parse(str)
		}
		return data, nil
	}
}

// stringToTimeHookFunc type conversion for string to time.Time
func stringToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String || t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		str, ok := data.(string)
		if !ok {
			return data, nil
		}

		return time.Parse(time.RFC3339, str)
	}
}
