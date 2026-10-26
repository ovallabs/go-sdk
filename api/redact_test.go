package api

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/ovalfi/go-sdk/model"
)

func newLoggedCall(baseURL string, buf *bytes.Buffer) *Call {
	return &Call{
		baseURL: baseURL + "/",
		client:  resty.New(),
		logger:  zerolog.New(buf),
	}
}

func Test_KYCCalls_DoNotLogResponseBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"status":200,"error":null,"data":{"id":"kyc-1","name":"Jane Doe","status":"verified",
			"biometrics_document_details":{"document_number":"X1234567"},"url":"https://verify.example/session-secret"}}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	tests := map[string]func(c *Call) error{
		"GetKYCByCustomerID": func(c *Call) error {
			_, err := c.GetKYCByCustomerID(context.Background(), "customer-1")
			return err
		},
		"GetVerifyCustomerKYC": func(c *Call) error {
			_, err := c.GetVerifyCustomerKYC(context.Background(), "customer-1", nil, nil)
			return err
		},
		"GetVerifyBiometricsLink": func(c *Call) error {
			_, _ = c.GetVerifyBiometricsLink(context.Background(), "customer-1")
			return nil
		},
	}

	for name, call := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			_ = call(newLoggedCall(ts.URL, &buf))

			logs := buf.String()
			require.Contains(t, logs, model.LogMsgResponseRedacted)
			require.Contains(t, logs, "/v1/kycs/customer-1")
			for _, secret := range []string{"Jane Doe", "X1234567", "session-secret"} {
				require.NotContains(t, logs, secret)
			}
		})
	}
}

func Test_VerifyCustomerKYC_MasksIDNumberInLogsAndErrors(t *testing.T) {
	const ssn = "123-45-6789"

	t.Run("success path logs the masked path", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// the real ID number must still reach the server
			require.Equal(t, "/v1/kycs/customer-1/ssn/"+ssn, r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(`{"status":200,"error":null,"data":{"verified":true}}`))
			require.NoError(t, err)
		}))
		defer ts.Close()

		var buf bytes.Buffer
		_, err := newLoggedCall(ts.URL, &buf).VerifyCustomerKYC(context.Background(), "customer-1", ssn, "ssn")
		require.NoError(t, err)

		require.NotContains(t, buf.String(), ssn)
		require.Contains(t, buf.String(), "/v1/kycs/customer-1/ssn/***")
	})

	t.Run("transport error does not leak the id number", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		url := ts.URL
		ts.Close() // connection refused, and net/http puts the full URL in the error

		var buf bytes.Buffer
		_, err := newLoggedCall(url, &buf).VerifyCustomerKYC(context.Background(), "customer-1", ssn, "ssn")

		require.Error(t, err)
		require.NotContains(t, err.Error(), ssn)
		require.NotContains(t, buf.String(), ssn)
	})
}

func Test_SubmitCustomerKYCDocument_DoesNotLogFormData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"status":200,"error":null,"data":{"id":"kyc-1","name":"Jane Doe"}}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	f, err := os.CreateTemp(t.TempDir(), "passport-jane-doe-*.jpg")
	require.NoError(t, err)
	defer f.Close()

	var buf bytes.Buffer
	_, err = newLoggedCall(ts.URL, &buf).SubmitCustomerKYCDocument(context.Background(), "customer-1", f, nil, "passport", "US")
	require.NoError(t, err)

	logs := buf.String()
	require.Contains(t, logs, model.LogMsgFormRedacted)
	require.NotContains(t, logs, "passport-jane-doe")
	require.NotContains(t, logs, "Jane Doe")
}

// Non-KYC calls keep logging exactly as before, so debugging elsewhere is unaffected.
func Test_NonKYCCalls_StillLogResponseBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"status":200,"error":null,"data":{"id":"cust-1","first_name":"Visible"}}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	var buf bytes.Buffer
	_, _ = newLoggedCall(ts.URL, &buf).GetCustomerByID(context.Background(), "cust-1")

	require.Contains(t, buf.String(), "Visible")
	require.NotContains(t, buf.String(), model.LogMsgResponseRedacted)
}
