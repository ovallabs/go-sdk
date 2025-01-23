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

	//deposit, err := apiCalls.InitiateDeposit(ctx, example.NewInitiateDepositRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Deposit: %+v\n", deposit)

	//deposits, err := apiCalls.GetAllDeposits(ctx, nil)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Deposits: %+v\n", deposits)

	//deposit, err := apiCalls.GetDepositID(ctx, "82f6e5b7-ad81-4ed5-bdac-153255b6aa17")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Deposit: %+v\n", deposit)

	//deposit, err := apiCalls.InternalFundsTransfer(ctx, example.NewFundTransferRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("deposit", deposit)

	//intraTransfer, err := apiCalls.IntraTransfer(ctx, example.NewIntraTransferRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Intra Transfer: ", intraTransfer)

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

	//transferResponse, err := apiCalls.InitiateTerminalTransfer(ctx, example.NewInitiateTerminalTransferRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfer response: ", transferResponse)

	//getTransfersResponse, err := apiCalls.GetTerminalTransfers(ctx, "", "USD", "NGN", nil, &model.Page{
	//	Number: helpers.GetPointerInt(1),
	//	Size:   helpers.GetPointerInt(5),
	//})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfers: ", getTransfersResponse)

	//transfer, err := apiCalls.GetTerminalTransferByID(ctx, "50a16aaa-1360-4423-a02e-6469c902ff17")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfer: ", transfer)

	//transferResponse, err := apiCalls.InitiateTransfer(ctx, example.NewInitiateTransferRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfer response: ", transferResponse)

	//rate, err := apiCalls.GetExchangeRates(ctx, 1000000, "NGN", "USD")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Rate: ", rate)

	//transfer, err := apiCalls.GetTransferByID(ctx, "7239478d-a6b7-40ee-85de-8b0a317c3771")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfer: ", transfer)

	//err := apiCalls.DeleteTransfer(ctx, "7239478d-a6b7-40ee-85de-8b0a317c3771", "Some reason")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transfer has been successfully deleted")

	//beneficiary, err := apiCalls.CreateBeneficiary(ctx, example.NewCreateBeneficiaryRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Beneficiary: %+v\n", beneficiary)

	//beneficiaries, err := apiCalls.GetBeneficiaries(ctx, "NGN", &model.Page{
	//	Number: helpers.GetPointerInt(1),
	//	Size:   helpers.GetPointerInt(5),
	//})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Beneficiaries: %+v\n", beneficiaries)

	//beneficiary, err := apiCalls.GetBeneficiaryByID(ctx, "c4158d8c-87a0-4f1b-b559-1aa2defd8495")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Printf("Beneficiary: %+v\n", beneficiary)

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

	//payout, err := apiCalls.GetPayoutByID(ctx, "9876e72e-e193-49cc-8c3a-eeab88476beb")
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

	//balances, err := apiCalls.GetCustomerBalances(ctx, "c4b9197f-009e-4019-b0dd-0cab6e9e3189")
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

	//transactions, err := apiCalls.GetTransactions(ctx,
	//	"c4b9197f-009e-4019-b0dd-0cab6e9e3189",
	//	"",
	//	"",
	//	"",
	//	"",
	//	nil,
	//	nil,
	//	&model.Page{
	//		Number: helpers.GetPointerInt(1),
	//		Size:   helpers.GetPointerInt(5),
	//	})
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transactions: ", transactions)

	//err := apiCalls.CancelTransaction(ctx, "e1a4b9a0-0c10-4842-809e-3acc8bca33b6", "transfer", "Some reason")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Transaction successfully cancelled")

	//balances, err := apiCalls.GetBalances(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Balances: ", balances)

	//account, err := apiCalls.GenerateBankAccount(ctx, example.NewGenerateBankAccountRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Account: ", account)

	//account, err := apiCalls.GetBankAccount(ctx, "c4b9197f-009e-4019-b0dd-0cab6e9e3189", "NGN")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Account: ", account)

	//tos, err := apiCalls.GetTermsOfService(context.Background(), "c4b9197f-009e-4019-b0dd-0cab6e9e3189", "USD")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Terms of Service: ", tos)

	//banks, err := apiCalls.GetBanks(ctx)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Banks", banks)

	//currency := "GHS"
	//country := "GH"
	//payoutType := "banks"
	//banks, err := apiCalls.GetSupportedBanks(ctx, currency, &country, &payoutType)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Banks", banks)

	//account, err := apiCalls.ResolveBankAccount(ctx, model.AccountResolveRequest{
	//	BankCode:      "044",
	//	AccountNumber: "9036678078",
	//})
	//
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Account: ", account)

	//err := apiCalls.MockDeposit(ctx, example.NewMockCustomerDepositRequest)
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//	return
	//}
	//fmt.Println("Deposit successful")
}
