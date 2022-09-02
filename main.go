package main

import (
	"os"

	"github.com/go-resty/resty/v2"
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
	apiCalls := api.New(&logger, client, model.PublicKey, model.BearerToken, model.BaseURL)
	apiCalls.RunInSandboxMode() // to ensure it is running in sandbox mode
	//ctx := context.Background()

	//portfolios, err := apiCalls.GetBusinessPortfolios(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("portfolios: %+v\n", portfolios)
	//"portfolio_id": "c7115f87-11aa-4d69-bcb4-c12dd7f5bf2f"

	//newCustomer, err := apiCalls.CreateCustomer(ctx, example.NewCreateCustomerRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new customer: %+v\n", newCustomer)
	//"id": "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd"

	//updatedCustomer, err := apiCalls.UpdateCustomer(ctx, example.NewUpdateCustomerRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new customer: %+v\n", updatedCustomer)

	//retrievedCustomer, err := apiCalls.GetCustomerByID(ctx, example.NewGetCustomerByIDRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new customer: %+v\n", retrievedCustomer)

	//retrievedCustomers, err := apiCalls.GetAllCustomers(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new customer: %+v\n", retrievedCustomers)

	//newYieldOffering, err := apiCalls.CreateYieldOfferingProfile(ctx, example.NewCreateYieldOfferingProfilesRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new yield offering: %+v\n", newYieldOffering)
	//"yield_offering_id": "ef8891af-e887-4e2c-ac79-7a9682d1ad77"
}
