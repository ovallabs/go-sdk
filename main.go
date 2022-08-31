package main

import (
	"os"

	"github.com/go-resty/resty"
	"github.com/rs/zerolog"

	"github.com/ovalfi/go-sdk/api"
	"github.com/ovalfi/go-sdk/model"
)

func main() {
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app is starting")
	defer logger.Info().Msg("stopped")
	apiCalls := api.New(&logger, client, model.Signature, model.BearerToken, model.BaseURL)
	apiCalls.RunInSandboxMode() // to ensure it is running in sandbox mode
}
