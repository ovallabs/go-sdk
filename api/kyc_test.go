package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// Test_GetKYCByCustomerID_DecodesFlatCustomerKYCDetails is a regression test
// for a bug where GetKYCByCustomerID always failed to decode a real
// response: external-api's GenericResponse envelope is only one layer deep
func Test_GetKYCByCustomerID_DecodesFlatCustomerKYCDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/v1/kycs/customer-123", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"status": 200,
			"message": "Customer KYC details retrieved successfully!",
			"error": null,
			"data": {
				"id": "kyc-1",
				"business_id": "biz-1",
				"customer_id": "customer-123",
				"kyc_provider": "sumsub",
				"status": "verified",
				"name": "Jane Doe",
				"biometrics_verified": true,
				"biometrics_document_details": {
					"first_name": "Jane",
					"last_name": "Doe",
					"document_type": "passport",
					"document_number": "X1234567",
					"selfie_url": "https://example.com/selfie.jpg",
					"video_url": "https://example.com/video.mp4"
				},
				"documents": []
			}
		}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	c := &Call{
		baseURL: ts.URL + "/",
		client:  resty.New(),
		logger:  zerolog.Nop(),
	}

	response, err := c.GetKYCByCustomerID(context.Background(), "customer-123")
	require.NoError(t, err)

	require.Equal(t, "kyc-1", response.Data.ID)
	require.Equal(t, "biz-1", response.Data.BusinessID)
	require.Equal(t, "customer-123", response.Data.CustomerID)
	require.Equal(t, "sumsub", response.Data.KYCProvider)
	require.Equal(t, "verified", response.Data.Status)
	require.Equal(t, "Jane Doe", response.Data.Name)
	require.True(t, response.Data.BiometricsVerified)
	require.NotNil(t, response.Data.BiometricsDocumentDetails)

	details, ok := response.Data.BiometricsDocumentDetails.(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "Jane", details["first_name"])
	require.Equal(t, "X1234567", details["document_number"])
	require.Equal(t, "https://example.com/video.mp4", details["video_url"])
}

// Test_GetKYCByCustomerID_EmptyDataFallsThroughCleanly confirms a response
// with no biometrics data yet
func Test_GetKYCByCustomerID_EmptyDataFallsThroughCleanly(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"status": 200,
			"message": "Customer KYC details retrieved successfully!",
			"error": null,
			"data": {
				"id": "kyc-2",
				"business_id": "biz-1",
				"customer_id": "customer-456",
				"status": "pending",
				"documents": []
			}
		}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	c := &Call{
		baseURL: ts.URL + "/",
		client:  resty.New(),
		logger:  zerolog.Nop(),
	}

	response, err := c.GetKYCByCustomerID(context.Background(), "customer-456")
	require.NoError(t, err)
	require.Equal(t, "pending", response.Data.Status)
	require.Nil(t, response.Data.BiometricsDocumentDetails)
}

// Test_SubmitCustomerKYCDocument_DecodesFlatCustomerKYCDetails is a
// regression test for the same envelope-depth bug GetKYCByCustomerID had:
// this endpoint also returns a flat KYC object (with its own string
// "status" field) nested one level under "data", not a second envelope.
// Targeting &response instead of &response.Data always failed to decode.
func Test_SubmitCustomerKYCDocument_DecodesFlatCustomerKYCDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/v1/kycs/customer-123/document", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"status": 200,
			"message": "Document submitted successfully!",
			"error": null,
			"data": {
				"id": "kyc-1",
				"business_id": "biz-1",
				"customer_id": "customer-123",
				"status": "pending",
				"documents": []
			}
		}`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	c := &Call{
		baseURL: ts.URL + "/",
		client:  resty.New(),
		logger:  zerolog.Nop(),
	}

	front, err := os.CreateTemp(t.TempDir(), "front-*.jpg")
	require.NoError(t, err)
	defer front.Close()

	response, err := c.SubmitCustomerKYCDocument(context.Background(), "customer-123", front, nil, "passport", "NG")
	require.NoError(t, err)

	require.Equal(t, "kyc-1", response.Data.ID)
	require.Equal(t, "biz-1", response.Data.BusinessID)
	require.Equal(t, "customer-123", response.Data.CustomerID)
	require.Equal(t, "pending", response.Data.Status)
}
