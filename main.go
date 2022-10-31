package main

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/ovalfi/go-sdk/api"
	"github.com/ovalfi/go-sdk/model"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app is starting")
	defer logger.Info().Msg("stopped")
	apiCalls := api.New(&logger, client, model.PublicKey, model.BearerToken, model.BaseURL)
	apiCalls.RunInSandboxMode() // to ensure it is running in sandbox mode
	//ctx := context.Background()

	/*newCustomer, err := apiCalls.CreateCustomer(ctx, example.NewCreateCustomerRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new customer: %+v\n", newCustomer)*/
	//"id": "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd"

	/*updatedCustomer, err := apiCalls.UpdateCustomer(ctx, example.NewUpdateCustomerRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new customer: %+v\n", updatedCustomer)*/

	/*retrievedCustomer, err := apiCalls.GetCustomerByID(ctx, example.NewGetCustomerByIDRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new customer: %+v\n", retrievedCustomer)

	retrievedCustomers, err := apiCalls.GetAllCustomers(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new customer: %+v\n", retrievedCustomers)*/

	//portfolios, err := apiCalls.GetBusinessPortfolios(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("portfolios: %+v\n", portfolios)
	//"portfolio_id": "c7115f87-11aa-4d69-bcb4-c12dd7f5bf2f"

	//newYieldOffering, err := apiCalls.CreateYieldOfferingProfile(ctx, example.NewCreateYieldOfferingProfilesRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new yield offering: %+v\n", newYieldOffering)
	//"yield_offering_id": "ef8891af-e887-4e2c-ac79-7a9682d1ad77"

	//updatedYieldOffering, err := apiCalls.UpdateYieldOfferingProfile(ctx, example.NewUpdateYieldOfferingProfilesRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("updated yield offering: %+v\n", updatedYieldOffering)
	//"yield_offering_id": "ef8891af-e887-4e2c-ac79-7a9682d1ad77"

	//yieldProfiles, err := apiCalls.GetAllYieldProfiles(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("yield profiles: %+v\n", yieldProfiles)

	//retrievedYieldProfile, err := apiCalls.GetYieldProfileByID(ctx, example.NewGetYieldProfileByIDRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("retrieved yield profile: %+v\n", retrievedYieldProfile)

	//newDeposit, err := apiCalls.InitiateDeposit(ctx, example.NewDepositRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new deposit initiated: %+v\n", newDeposit)

	//deposits, err := apiCalls.GetAllDeposits(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("get all deposits initiated: %+v\n", deposits)

	//batchDeposit, err := apiCalls.GetDepositByBatchID(ctx, "2022-09-05T00:00:00Z")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("get deposit by batchDateID initiated: %+v\n", batchDeposit)

	//newWithdrawal, err := apiCalls.InitiateWithdrawal(ctx, example.NewInitiateWithdrawalRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("new withdrawal initiated: %+v\n", newWithdrawal)

	/*wallet, err := apiCalls.GetWallet(ctx, model.WalletRequest{
		CustomerID: "bb1f2b22-0b5c-4c1c-a8d1-df99f02e08de",
		Network:    "TEST",
		Asset:      "USDC",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("wallet info: %+v\n", wallet)

	allWallet, err := apiCalls.GetWallets(ctx, "bb1f2b22-0b5c-4c1c-a8d1-df99f02e08de")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("all wallet info: %+v\n", *allWallet[0])*/
	/*assets, err := apiCalls.GetSupportedAssets(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("assets", assets)
	balances, err := apiCalls.GetCustomerBalances(ctx, uuid.MustParse("cefec56e-3781-4b3a-bda6-ba4e7c0e49cd"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("balances", balances)

	balance, err := apiCalls.GetCustomerBalance(ctx, model.GetCustomerBalanceRequest{
		CustomerID:      uuid.MustParse("cefec56e-3781-4b3a-bda6-ba4e7c0e49cd"),
		YieldOfferingID: uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77"),
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("balances", balance)
	deposit, err := apiCalls.InternalFundsTransfer(ctx, model.FundTransferRequest{
		CustomerID:      uuid.MustParse("cefec56e-3781-4b3a-bda6-ba4e7c0e49cd"),
		Reference:       "ddffd",
		Amount:          10,
		Action:          model.Credit,
		YieldOfferingID: uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77"),
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("deposit", deposit)*/
	/*cID := uuid.MustParse("5e37dc39-5b70-492a-a5ad-46c75d06111e")
	size := 1
	transaction, err := apiCalls.GetTransactions(ctx, &model.TransactionRequest{CustomerID: &cID, Size: &size})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("transaction", transaction)*/
	/*deposit, err := apiCalls.GetDepositID(ctx, uuid.MustParse("9c6c34d9-49b1-47c6-88f6-98ca0163c597"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("deposit", deposit)*/

	/*details, err := apiCalls.GetExchangeRates(context.Background(), model.GetExchangeRateRequest{
		Amount:              3000,
		SourceCurrency:      "USD",
		DestinationCurrency: "NGN",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("details", details)*/

	/*banks, err := apiCalls.GetBanks(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("banks", banks)*/

	/*account, err := apiCalls.ResolveBankAccount(context.Background(), model.AccountResolveRequest{
		BankCode:      "057",
		AccountNumber: "2209276822",
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("account", account)*/
	/*withdrawal, err := apiCalls.FiatWithdrawal(ctx, model.WithdrawalRequest{
		CustomerID:      uuid.MustParse("9f40fb69-64e3-4d23-853a-0243af155427"),
		Reference:       "polkj",
		Amount:          10,
		YieldOfferingID: uuid.MustParse("42ee80d8-2a95-419c-aad1-5643d306948e"),
		PayoutCurrency:  "NGN",
		BankDetail: &model.BankDetail{
			BankCode:      "057",
			AccountNumber: "2209276822",
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("withdrawal", withdrawal)*/
}
