package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/ovalfi/go-sdk/model"
)

// Test_UpdateCustomerPaymentCard sends the update to the card's path and decodes the updated card
func Test_UpdateCustomerPaymentCard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		require.Equal(t, "/v1/payments/cards/customer-123/card-456", r.URL.Path)

		var body model.UpdatePaymentCardRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		require.Equal(t, model.UpdatePaymentCardRequest{Status: "approved", Reason: "address confirmed", UpdatedBy: "admin@torus.test"}, body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"status": 200,
			"message": "Updated payment card successfully!",
			"error": null,
			"data": {"id": "6f1c3b7e-2a4d-4c8e-9b1a-0d2e3f4a5b6c", "status": "approved", "last_four_digits": "1234"}
		}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	c := &Call{
		baseURL: ts.URL + "/",
		client:  resty.New(),
		logger:  zerolog.Nop(),
	}

	card, err := c.UpdateCustomerPaymentCard(context.Background(), "customer-123", "card-456", model.UpdatePaymentCardRequest{
		Status:    "approved",
		Reason:    "address confirmed",
		UpdatedBy: "admin@torus.test",
	})
	require.NoError(t, err)
	require.Equal(t, "approved", card.Status)
	require.Equal(t, "1234", card.LastFourDigits)
}
