package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/ovalfi/go-sdk/model"
)

func TestGetCompetitorsRates(t *testing.T) {
	expected := []model.CompetitorRate{
		{
			Provider:     "remitly",
			ProviderName: "Remitly",
			From:         "USD",
			To:           "NGN",
			Rate:         1500.25,
		},
		{
			Provider:     "nala",
			ProviderName: "Nala",
			From:         "USD",
			To:           "NGN",
			Rate:         1498.5,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/utils/competitors-rates", r.URL.Path)
		assert.Equal(t, "USD", r.URL.Query().Get("from"))
		assert.Equal(t, "NGN", r.URL.Query().Get("to"))

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := &Call{
		baseURL:     ts.URL + "/",
		client:      resty.New(),
		logger:      zerolog.Nop(),
		bearerToken: "test-token",
	}

	rates, err := call.GetCompetitorsRates(context.Background(), "USD", "NGN")
	assert.NoError(t, err)
	assert.Equal(t, expected, rates)
}
