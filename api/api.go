// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/ovalfi/go-sdk/model"
)

// RemoteCalls abstracted definition of supported functions
type RemoteCalls interface {
	// Customer APIs
	CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (model.Customer, error)
	UpdateCustomer(ctx context.Context, request model.UpdateCustomerRequest) (model.Customer, error)
	GetAllCustomers(ctx context.Context) (model.AllCustomersResponse, error)
	GetCustomerByID(ctx context.Context, customerID string) (model.Customer, error)
	GetCustomerBalance(ctx context.Context, customerID, yieldOfferingID string) (model.CustomerBalance, error)
	GetCustomerBalances(ctx context.Context, customerID string) (model.CustomerBalances, error)
	DeleteCustomer(ctx context.Context, customerID string) error

	// Transfer API
	InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.TransferResponse, error)
	GetExchangeRates(ctx context.Context, amount float64, sourceCurrency, destinationCurrency string) (model.ExchangeRateDetails, error)
	GetTransferByID(ctx context.Context, transferID string) (model.Transfer, error)
	DeleteTransfer(ctx context.Context, transferID, reason string) error
	DeleteTransferBatch(ctx context.Context, batchDate, currency, reason string) error
	InitiateTerminalTransfer(ctx context.Context, request model.InitiateTerminalTransferRequest) (model.TerminalTransfer, error)
	GetTerminalTransfers(ctx context.Context, status, sourceCurrency, destinationCurrency string, dateBetween *model.DateBetween, page *model.Page) (model.AllTransfersResponse, error)
	GetTerminalTransferByID(ctx context.Context, transferID string) (model.TerminalTransfer, error)
	GetSettlementByID(ctx context.Context, settlementID string) (model.Settlement, error)

	// Transaction APIs
	GetTransactions(ctx context.Context, customerID, yieldOfferingID, status, reference, batchDate string, amount *float64, dateBetween *model.DateBetween, page *model.Page) (model.AllTransactionsResponse, error)
	CancelTransaction(ctx context.Context, transactionID, transactionType, reason string) error
	CancelBatchTransaction(ctx context.Context, batchDate, transactionType, currency, reason string) error
	GetBalances(ctx context.Context) (map[string]float64, error)

	// Payment APIs
	GetBanks(ctx context.Context) ([]model.BankCode, error)
	GetSupportedBanks(ctx context.Context, currency, country string) ([]model.Bank, error)
	ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetails, error)
	GenerateBankAccount(ctx context.Context, request model.GenerateBankAccountRequest) (model.BankAccount, error)
	GetBankAccount(ctx context.Context, customerID, currency string) (model.BankAccount, error)
	MockDeposit(ctx context.Context, request model.MockCustomerDepositRequest) error

	// Payout APIs
	GetPayoutByID(ctx context.Context, payoutID string) (model.PayoutResponse, error)
	InitiateDirectBulkPayout(ctx context.Context, request model.InitiateBulkPayoutRequest) (model.PayoutDetails, error)
	InitiatePayout(ctx context.Context, currency, payoutType, beneficiaryType, remarks string, customerID *string, document *os.File) (model.PayoutDetails, error)
	GetAllPayouts(ctx context.Context, status, search string, dateBetween model.DateBetween, page model.Page) (model.AllPayoutsResponse, error)
	CancelPayout(ctx context.Context, request model.CancelPayoutRequest) error
	UpdatePayoutAccount(ctx context.Context, payoutID string, request model.TransferBeneficiaryDetails) error
	GetPayoutConfig(ctx context.Context, currency string) (model.BulkPayoutConfig, error)
	GetPayoutDocumentTemplate(ctx context.Context, currency, docType string) (string, error)

	// Currency Swap APIs
	InitiateCurrencySwap(ctx context.Context, request model.InitiateCurrencySwapRequest) (model.CurrencySwap, error)
	GetCurrencySwaps(ctx context.Context, status, from, to string, dateBetween *model.DateBetween, page *model.Page) (model.AllSwapsResponse, error)
	GetCurrencySwapByID(ctx context.Context, currencySwapID string) (model.CurrencySwap, error)

	// Beneficiary APIs
	CreateBeneficiary(ctx context.Context, request model.CreateBeneficiaryRequest) (model.TransferBeneficiary, error)
	GetBeneficiaries(ctx context.Context, currency string, page *model.Page) (model.AllBeneficiariesResponse, error)
	GetBeneficiaryByID(ctx context.Context, beneficiaryID string) (model.TransferBeneficiary, error)

	// Deposit APIs
	InitiateDeposit(ctx context.Context, request model.InitiateDepositRequest) (model.Deposit, error)
	GetAllDeposits(ctx context.Context, settled *bool) (model.DepositBatchResponse, error)
	GetDepositID(ctx context.Context, id string) (model.Deposit, error)
	InternalFundsTransfer(ctx context.Context, request model.FundTransferRequest) (model.Deposit, error)
	IntraTransfer(ctx context.Context, request model.IntraTransferRequest) (model.IntraTransferResponse, error)

	// Withdrawal APIs
	InitiateWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error)
	FiatWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error)
	CryptoWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error)
	FeeWithdrawal(ctx context.Context, request model.FeeWithdrawalRequest) (model.FeeWithdrawalResponse, error)

	// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
	RunInSandboxMode()
}

// Call object
type Call struct {
	client       *resty.Client
	logger       zerolog.Logger
	baseURL      string
	publicKey    string
	bearerToken  string
	sandboxMode  bool
	idempotentID uuid.UUID
}

// New initialises the object Call
func New(z *zerolog.Logger, c *resty.Client, publicKey, bearerToken, bURL string) RemoteCalls {
	call := &Call{
		client:       c,
		logger:       z.With().Str("sdk", "ovalfi").Logger(),
		baseURL:      bURL,
		publicKey:    publicKey,
		bearerToken:  bearerToken,
		idempotentID: uuid.New(),
	}
	return RemoteCalls(call)
}

// ReloadIdempotentID this reissues a new idempotent ID
func (c *Call) ReloadIdempotentID() {
	c.idempotentID = uuid.New()
}

// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
func (c *Call) RunInSandboxMode() {
	c.client.SetDebug(true)
	c.sandboxMode = true
}
