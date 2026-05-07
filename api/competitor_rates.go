package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

// GetCompetitorsRates fetches the latest competitor rates for a currency pair.
func (c *Call) GetCompetitorsRates(ctx context.Context, from, to string) ([]model.CompetitorRate, error) {
	var (
		err      error
		response []model.CompetitorRate
		params   = map[string]interface{}{
			"from": from,
			"to":   to,
		}
		path = fmt.Sprintf("%s/competitors-rates", utilAPIVersion)
	)

	err = c.makeRequest(ctx, path, http.MethodGet, nil, params, nil, nil, &response)
	return response, err
}
