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

	//currencySwap, err := apiCalls.InitiateCurrencySwap(ctx, example.NewInitiateCurrencySwapRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Currency swap: %+v\n", currencySwap)

	//swaps, err := apiCalls.GetCurrencySwaps(ctx, "completed", "USD", "NGN", nil, &model.Page{
	//	Number: helpers.GetPointerInt(1),
	//	Size:   helpers.GetPointerInt(5),
	//})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Swaps: %+v\n", swaps)

	//swap, err := apiCalls.GetCurrencySwapByID(ctx, "687cf078-c553-47be-b99c-708a6abc9a44")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Swap: %+v\n", swap)

	//doc, err := apiCalls.GetPayoutDocumentTemplate(ctx, "USD", "banks")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Document template: %+v\n", doc)

	//payoutDetails, err := apiCalls.InitiateDirectBulkPayout(ctx, example.NewInitiateBulkPayoutRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Payout details: %+v\n", payoutDetails)

	//file, err := os.Open("/Users/z/Downloads/Document_Oval.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//payoutDetails, err := apiCalls.InitiatePayout(ctx, "USD", "banks", string(model.MultiplePayout), "Some remarks", file)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Payout details: %+v\n", payoutDetails)

	//payouts, err := apiCalls.GetAllPayouts(ctx, "pending", "", model.DateBetween{}, model.Page{
	//	Number: helpers.GetPointerInt(1),
	//	Size:   helpers.GetPointerInt(5),
	//})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Payouts: %+v\n", payouts)

	//config, err := apiCalls.GetPayoutConfig(ctx, "USD")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Payout config: %+v\n", config)

	//payout, err := apiCalls.GetPayoutByID(ctx, "ef467f44-ed91-4875-8861-c2a5c7e4232d")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Payout: %+v\n", payout)

	//isCancelled, err := apiCalls.CancelPayout(ctx, example.NewCancelPayoutRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Has payout been cancelled: %+v\n", isCancelled)

	//customer, err := apiCalls.CreateCustomer(ctx, example.NewCreateCustomerRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Customer: %+v\n", customer)

	/*updatedCustomer, err := apiCalls.UpdateCustomer(ctx, example.NewUpdateCustomerRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new customer: %+v\n", updatedCustomer)*/

	//customer, err := apiCalls.GetCustomerByID(ctx, "2d0378b6-a707-41ec-8636-6b3900ef60fd")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Customer: %+v\n", customer)

	//customers, err := apiCalls.GetAllCustomers(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Customers: %+v\n", customers)

	//balance, err := apiCalls.GetCustomerBalance(ctx, "625473e1-1dbf-446c-b86d-005d5eae0919", "21c48a42-d840-4f66-bdb0-c7510a038bd4")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Balance: %+v\n", balance)

	//balances, err := apiCalls.GetCustomerBalances(ctx, "625473e1-1dbf-446c-b86d-005d5eae0919")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Balances: %+v\n", balances)

	//err := apiCalls.DeleteCustomer(ctx, "6cef5231-fc1e-45b3-a9ae-4d204245b0ae")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Customer deleted successfully")

	/*portfolios, err := apiCalls.GetBusinessPortfolios(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("portfolios: %+v\n", portfolios)
	"portfolio_id": "c7115f87-11aa-4d69-bcb4-c12dd7f5bf2f"*/

	/*newYieldOffering, err := apiCalls.CreateYieldOfferingProfile(ctx, example.NewCreateYieldOfferingProfilesRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("new yield offering: %+v\n", newYieldOffering)
	"yield_offering_id": "ef8891af-e887-4e2c-ac79-7a9682d1ad77"*/

	/*updatedYieldOffering, err := apiCalls.UpdateYieldOfferingProfile(ctx, example.NewUpdateYieldOfferingProfilesRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("updated yield offering: %+v\n", updatedYieldOffering)
	"yield_offering_id": "ef8891af-e887-4e2c-ac79-7a9682d1ad77"*/

	/*yieldProfiles, err := apiCalls.GetAllYieldProfiles(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("yield profiles: %+v\n", yieldProfiles)*/

	/*retrievedYieldProfile, err := apiCalls.GetYieldProfileByID(context.Background(), example.NewGetYieldProfileByIDRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("retrieved yield profile: %+v\n", retrievedYieldProfile)*/

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
		Asset:      "USDT",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("wallet info: %+v\n", wallet)

	/*allWallet, err := apiCalls.GetWallets(ctx, "bb1f2b22-0b5c-4c1c-a8d1-df99f02e08de")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("all wallet info: %+v\n", *allWallet[0])
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

	/*intraTransfer, err := apiCalls.IntraTransfer(ctx, model.IntraTransferRequest{
		Reference: "some-test-reference-01",
		Amount:    10,
		Sender: model.TransferParty{
			CustomerID:      "4a31a43b-7c54-4578-a020-87d1d2b0f6f5",
			YieldOfferingID: "21c48a42-d840-4f66-bdb0-c7510a038bd4",
		},
		Receiver: model.TransferParty{
			CustomerID:      "a6c04a14-9214-4cd1-9945-cb7a04b9fa07",
			YieldOfferingID: "d0f25fb9-b3ef-42ba-b1f9-e946722b064f",
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("intra-transfer", intraTransfer)*/
	/*cID := uuid.MustParse("5e37dc39-5b70-492a-a5ad-46c75d06111e")
	size := 1
	transaction, err := apiCalls.GetTransactions(ctx, &model.TransactionRequest{CustomerID: &cID, Size: &size})
	//cID := uuid.MustParse("5e37dc39-5b70-492a-a5ad-46c75d06111e")
	//size := 1
	/*ref := "07b1057f-e85e-441e-ac4f-d1a498617762"
	transaction, err := apiCalls.GetTransactions(context.Background(), &model.TransactionRequest{Reference: &ref})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("transaction", *transaction.Items.Transactions[0])*/
	/*deposit, err := apiCalls.GetDepositID(ctx, uuid.MustParse("9c6c34d9-49b1-47c6-88f6-98ca0163c597"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("deposit", deposit)*/

	/*details, err := apiCalls.GetExchangeRates(context.Background(), model.GetExchangeRateRequest{
		Amount:              3000,
		SourceCurrency:      "USD",
		DestinationCurrency: "TRON",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("details", details)*/

	/*details, err := apiCalls.InitiateTransfer(context.Background(), model.InitiateTransferRequest{
		CustomerID:  "9f40fb69-64e3-4d23-853a-0243af155427",
		Amount:      3000,
		Currency:    "USD",
		Destination: model.TransferDestination{},
		Note:        "",
		Reason:      "",
		Reference:   "",
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
	/*account, err := apiCalls.GenerateBankAccount(context.Background(), model.BankAccountRequest{
		CustomerID:  uuid.MustParse("163c4dbb-7032-4ee7-8ab1-75b6cc01726f"),
		BVN:         "22481815830",
		Reference:   "ref5679",
		PhoneNumber: "2349090909091",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("account", account)

	bank, err := apiCalls.GetBankAccount(ctx, uuid.MustParse("163c4dbb-7032-4ee7-8ab1-75b6cc01726f"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("bank", bank)*/

	//feeWithdrawal, err := apiCalls.FeeWithdrawal(context.Background(), model.FeeWithdrawalRequest{
	//	CustomerID:          uuid.MustParse("4a31a43b-7c54-4578-a020-87d1d2b0f6f5"),
	//	BusinessID:          uuid.MustParse("b21a44b0-c25b-474b-a986-8af627109c19"),
	//	Reference:           "c982d536-7a1d-4034-bc1a-0f3527da1c28",
	//	WithdrawalReference: "dae68c86-d80b-4652-9bec-38f2daf4db86",
	//	FeeType:             "amount",
	//	Amount:              10,
	//	Percentage:          0,
	//	YieldOfferingID:     uuid.MustParse("21c48a42-d840-4f66-bdb0-c7510a038bd4"),
	//})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("withdrawal", feeWithdrawal)

	//err := apiCalls.DeleteCustomer(context.Background(), uuid.MustParse("260483ca-07b4-4ab6-a224-2ddf33144439"))
	//
	//if err != nil {
	//	fmt.Printf("Error: %v\n\n", err)
	//	return
	//} else {
	//	fmt.Printf("Succesfuly deleted")
	//}
}
