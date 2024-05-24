package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/ovalfi/go-sdk/model"
)

func Test_makeRequest(t *testing.T) {
	tests := map[string]struct {
		requestPath      string
		expectedResponse interface{}
		errorResponse    model.ErrorResponse
		expectedErr      error
		handlerFunc      http.HandlerFunc
	}{
		"Success response with params GET request": {
			requestPath: "/name",
			expectedResponse: struct {
				Data struct {
					Name string `json:"name"`
				} `json:"data"`
			}{
				Data: struct {
					Name string `json:"name"`
				}{
					Name: "John Doe",
				},
			},
			expectedErr: nil,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/name", r.URL.Path)
				query := r.URL.Query()
				firstName := query.Get("first_name")
				lastName := query.Get("last_name")
				body, err := json.Marshal(struct {
					Data struct {
						Name string `json:"name"`
					} `json:"data"`
				}{
					Data: struct {
						Name string `json:"name"`
					}{
						Name: func(firstName, lastName string) string {
							return fmt.Sprintf("%s %s", firstName, lastName)
						}(firstName, lastName),
					},
				})
				assert.NoError(t, err)

				w.WriteHeader(http.StatusOK)
				_, err = w.Write(body)
				assert.NoError(t, err)
			},
		},
		"Success response with request body POST request": {
			requestPath: "/register",
			expectedResponse: struct {
				Data struct {
					Message string `json:"message"`
				} `json:"data"`
			}{
				Data: struct {
					Message string `json:"message"`
				}{
					Message: "User registered successfully!",
				},
			},
			expectedErr: nil,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/register", r.URL.Path)

				s := struct {
					FirstName string `json:"first_name"`
					LastName  string `json:"last_name"`
				}{}

				err := json.NewDecoder(r.Body).Decode(&s)
				assert.NoError(t, err)

				body, err := json.Marshal(struct {
					Data struct {
						Message string `json:"message"`
					} `json:"data"`
				}{
					Data: struct {
						Message string `json:"message"`
					}{
						Message: "User registered successfully!",
					},
				})
				assert.NoError(t, err)

				w.WriteHeader(http.StatusOK)
				_, err = w.Write(body)
				assert.NoError(t, err)
			},
		},
		"Failed with error response": {
			requestPath: "/error",
			errorResponse: model.ErrorResponse{
				Message: "something went wrong",
			},
			expectedErr: errors.New("something went wrong"),
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/error", r.URL.Path)
				body, err := json.Marshal(model.ErrorResponse{
					Message: "something went wrong",
				})
				assert.NoError(t, err)

				w.WriteHeader(http.StatusUnauthorized)
				_, err = w.Write(body)
				assert.NoError(t, err)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			ts := httptest.NewServer(tt.handlerFunc)
			defer ts.Close()

			c := &Call{
				baseURL: ts.URL,
				client:  resty.New(),
				logger:  zerolog.Nop(),
			}
			if tt.requestPath == "/name" {
				var responseBody struct {
					Data struct {
						Name string `json:"name"`
					} `json:"data"`
				}
				params := map[string]interface{}{
					"first_name": "John",
					"last_name":  "Doe",
				}
				err := c.makeRequest(ctx, "/name", http.MethodGet, params, nil, nil, &responseBody)
				assert.Equal(t, tt.expectedResponse, responseBody)
				assert.Equal(t, tt.expectedErr, err)
			} else if tt.requestPath == "/register" {
				var responseBody struct {
					Data struct {
						Message string `json:"message"`
					} `json:"data"`
				}
				request := struct {
					FirstName string `json:"first_name"`
					LastName  string `json:"last_name"`
				}{
					FirstName: "John",
					LastName:  "Doe",
				}
				err := c.makeRequest(ctx, "/register", http.MethodPost, nil, nil, request, &responseBody)
				assert.Equal(t, tt.expectedResponse, responseBody)
				assert.Equal(t, tt.expectedErr, err)
			} else if tt.requestPath == "/error" {
				var responseBody struct{} // not needed anyway
				err := c.makeRequest(ctx, "/error", http.MethodGet, nil, nil, nil, &responseBody)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
