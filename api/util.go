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
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

// withRedactedLogging returns a context that makes makeRequest redact its logs for the call.
// maskedPath is the path to log in place of the real one, e.g. with an ID number replaced by ***.
func withRedactedLogging(ctx context.Context, maskedPath string) context.Context {
	return context.WithValue(ctx, model.RedactedLogContextKey, model.RedactedLog{MaskedPath: maskedPath})
}

func redactedLogFrom(ctx context.Context) (model.RedactedLog, bool) {
	r, ok := ctx.Value(model.RedactedLogContextKey).(model.RedactedLog)
	return r, ok
}

// maskedError carries a sanitized message while keeping the original error in the chain for errors.Is/As
type maskedError struct {
	msg   string
	cause error
}

func (e *maskedError) Error() string { return e.msg }
func (e *maskedError) Unwrap() error { return e.cause }

// maskEndpoint replaces the real endpoint in err's text and drops any query string that follows it,
// as net/http copies the full URL into transport errors. The original error stays reachable via Unwrap.
func maskEndpoint(err error, endpoint, maskedEndpoint string) error {
	if err == nil || !strings.Contains(err.Error(), endpoint) {
		return err
	}
	msg := strings.ReplaceAll(err.Error(), endpoint, maskedEndpoint)
	msg = regexp.MustCompile(regexp.QuoteMeta(maskedEndpoint)+`\?[^\s"']*`).ReplaceAllString(msg, maskedEndpoint)
	return &maskedError{msg: msg, cause: err}
}

func (c *Call) makeRequest(ctx context.Context, path, method string, signature *string, params, formData map[string]interface{}, requestBody, responseData interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	redact, redacted := redactedLogFrom(ctx)
	logEndpoint := endpoint
	if redacted {
		logEndpoint = fmt.Sprintf("%s%s", c.baseURL, redact.MaskedPath)
	}

	log := c.logger.With().Str("method", method).Str("endpoint", logEndpoint).Logger()
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

	if redacted {
		// resty's debug dump prints the real URL and bodies; turn it off for this request only
		client.SetDebug(false)
	}

	if signature != nil {
		client = client.SetHeader("Signature", *signature)
	}

	if requestBody != nil {
		if redacted {
			log.Info().Msg(model.LogMsgRequestRedacted)
		} else {
			log.Info().Interface(model.LogStrRequest, requestBody).Msg("request")
		}
		client.SetBody(requestBody)
	}

	if params != nil {
		if redacted {
			log.Info().Msg(model.LogMsgParamsRedacted)
		} else {
			log.Info().Interface(model.LogStrParams, params).Msg("parameters")
		}
		for k, v := range params {
			client.SetQueryParam(k, v.(string))
		}
	}

	if formData != nil {
		if redacted {
			log.Info().Msg(model.LogMsgFormRedacted)
		} else {
			log.Info().Interface(model.LogStrForm, formData).Msg("form data")
		}
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
	case http.MethodDelete:
		res, err = client.Delete(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return err
	}

	if err != nil {
		if redacted {
			err = maskEndpoint(err, endpoint, logEndpoint)
			log.Err(err).Msg("something went wrong")
			if res != nil {
				log.Err(err).Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(model.LogMsgErrorBodyRedacted)
			}
			return err
		}
		log.Err(err).Msg("something went wrong")
		if res != nil {
			log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		}
		return err
	}

	if genericResponse.Error != nil {
		err = errors.New(genericResponse.Error.Details)
		if redacted {
			// the details may echo submitted identity values; callers still get them in the returned error
			log.Error().Msg(model.LogMsgErrorBodyRedacted)
			return err
		}
		log.Err(err).Msg("error while making request")
		return err
	}

	if redacted {
		log.Info().Msg(model.LogMsgResponseRedacted)
	} else {
		log.Info().Interface(model.LogStrResponse, genericResponse.Data).Msg("response")
	}
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
			stringToTimeHookFunc(time.RFC3339Nano),
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
func stringToTimeHookFunc(layout string) mapstructure.DecodeHookFunc {
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

		return time.Parse(layout, str)
	}
}
