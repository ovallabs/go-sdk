package api

import (
	"context"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

const reportAPIVersion = "v1/transaction-reporting"

// SubmitSTR makes request to Torus to submit a Suspicious Transaction Report
func (c *Call) SubmitSTR(ctx context.Context, request model.SubmitSTRRequest) error {
	path := reportAPIVersion + "/str"

	return c.makeRequest(ctx, path, http.MethodPost, nil, nil, nil, request, nil)
}
