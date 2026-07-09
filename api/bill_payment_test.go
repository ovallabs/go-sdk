package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/ovalfi/go-sdk/model"
)

func newTestCall(url string) *Call {
	return &Call{
		baseURL:     url + "/",
		client:      resty.New(),
		logger:      zerolog.Nop(),
		bearerToken: "test-token",
	}
}

func TestGetBillerCategories(t *testing.T) {
	expected := []model.BillerCategory{
		{Code: "airtime", Name: "Airtime"},
		{Code: "electricity", Name: "Electricity"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/bills/NG/categories", r.URL.Path)

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	categories, err := call.GetBillerCategories(context.Background(), "NG")
	assert.NoError(t, err)
	assert.Equal(t, expected, categories)
}

func TestGetBillers(t *testing.T) {
	expected := []model.Biller{
		{Code: "mtn", Name: "MTN", PaymentTypes: []string{"prepaid"}},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/bills/NG/categories/airtime/billers", r.URL.Path)

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	billers, err := call.GetBillers(context.Background(), "airtime", "NG")
	assert.NoError(t, err)
	assert.Equal(t, expected, billers)
}

func TestGetBillerProducts(t *testing.T) {
	expected := model.AllBillerProductsResponse{
		Items: []model.BillerProduct{
			{Code: "ekedc-prepaid-1000", Name: "EKEDC Prepaid 1000 NGN", CategoryCode: "electricity", BillerCode: "ekedc", PaymentType: "prepaid"},
		},
		Page: model.PageInfo{Page: 1},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/bills/NG/categories/electricity/billers/ekedc/products", r.URL.Path)
		assert.Equal(t, "prepaid", r.URL.Query().Get("payment_type"))
		assert.Equal(t, "1", r.URL.Query().Get("number"))

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	paymentType := "prepaid"
	number := 1
	products, err := call.GetBillerProducts(context.Background(), "electricity", "ekedc", "NG", &paymentType, &model.Page{Number: &number})
	assert.NoError(t, err)
	assert.Equal(t, expected, products)
}

func TestValidateBillerCustomer(t *testing.T) {
	request := model.ValidateBillerCustomerRequest{Code: "ekedc-prepaid", CustomerID: "1234567890"}
	expected := model.ValidateBillerCustomerResponse{CustomerName: "John Doe", RequireValidationReference: true, ValidationReference: "ref-123"}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/bills/validate-customer", r.URL.Path)

		var received model.ValidateBillerCustomerRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&received))
		assert.Equal(t, request, received)

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	resp, err := call.ValidateBillerCustomer(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestPayBill(t *testing.T) {
	validationRef := "ref-123"
	request := model.PayBillRequest{Code: "ekedc-prepaid", CustomerID: "1234567890", Amount: 5000, ValidationReference: &validationRef}
	expected := model.PayBillResponse{ID: "bp-1", Status: "pending"}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/bills/pay", r.URL.Path)

		var received model.PayBillRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&received))
		assert.Equal(t, request, received)

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	resp, err := call.PayBill(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestGetBillPaymentTransaction(t *testing.T) {
	expected := model.BillPaymentTransaction{
		ID:         "bp-1",
		Code:       "ekedc-prepaid",
		CustomerID: "1234567890",
		Amount:     5000,
		Currency:   "NGN",
		Status:     "pending",
		CreatedAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/bills/payments/bp-1", r.URL.Path)

		body, err := json.Marshal(model.GenericResponse{Data: expected})
		assert.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	call := newTestCall(ts.URL)

	resp, err := call.GetBillPaymentTransaction(context.Background(), "bp-1")
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}
