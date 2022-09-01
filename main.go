package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/ovalfi/go-sdk/api"
	"github.com/ovalfi/go-sdk/model"
	"github.com/ovalfi/go-sdk/model/example"
)

func main() {
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app is starting")
	defer logger.Info().Msg("stopped")
	apiCalls := api.New(&logger, client, model.Signature, model.BearerToken, model.BaseURL)
	apiCalls.RunInSandboxMode() // to ensure it is running in sandbox mode
	ctx := context.Background()

	newCustomer, err := apiCalls.CreateCustomer(ctx, example.NewCustomerRequest)
	fmt.Printf("new customer: %+v\n", newCustomer)
	fmt.Printf("error: %+v\n", err)
}
