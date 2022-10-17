package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ovalfi/go-sdk/model"
)

const transactionAPIVersion = "v1/transaction"

// GetTransactions makes an API request using Call to get transactions
func (c *Call) GetTransactions(ctx context.Context, request *model.TransactionRequest) (model.TransactionResponse, error) {
	params := url.Values{}
	if request.CustomerID != nil {
		cID := *request.CustomerID
		params.Set("customer_id", cID.String())
	}

	if request.YieldOfferingID != nil {
		yID := *request.YieldOfferingID
		params.Set("yield_offering_id", yID.String())
	}

	if request.Reference != nil {
		params.Set("reference", *request.Reference)
	}

	if request.BatchDate != nil {
		params.Set("batch_date", *request.BatchDate)
	}

	if request.Size != nil {
		size := strconv.Itoa(*request.Size)
		params.Set("size", size)
	}

	if request.Page != nil {
		page := strconv.Itoa(*request.Page)
		params.Set("page", page)
	}

	var endpoint string
	if request == nil {
		endpoint = fmt.Sprintf("%s%s", c.baseURL, transactionAPIVersion)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", c.baseURL, transactionAPIVersion, params.Encode())
	}

	response := struct {
		Data model.TransactionResponse `json:"data"`
	}{}

	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return model.TransactionResponse{}, err
	}

	if res.StatusCode() != http.StatusOK {
		return model.TransactionResponse{}, model.ErrNetworkError
	}

	return response.Data, nil
}
