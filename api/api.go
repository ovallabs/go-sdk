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
	GetAllCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomerByID(ctx context.Context, request model.GetCustomerByIDRequest) (model.CustomerInfo, error)
	GetCustomerBalance(ctx context.Context, request model.GetCustomerBalanceRequest) (model.CustomerBalanceResponse, error)
	GetCustomerBalances(ctx context.Context, customerID uuid.UUID) (model.CustomerBalancesResponse, error)
	DeleteCustomer(ctx context.Context, customerID uuid.UUID) error

	// Yield APIs
	GetBusinessPortfolios(ctx context.Context) ([]model.Portfolio, error)
	CreateYieldOfferingProfile(ctx context.Context, request model.CreateYieldOfferingProfilesRequest) (model.YieldOfferingProfile, error)
	UpdateYieldOfferingProfile(ctx context.Context, request model.UpdateYieldOfferingProfilesRequest) (model.UpdatedYieldOfferingProfile, error)
	GetAllYieldProfiles(ctx context.Context) ([]model.YieldOfferingProfile, error)
	GetYieldProfileByID(ctx context.Context, request model.GetYieldProfileByIDRequest) (model.YieldOfferingProfile, error)

	// Deposit APIs
	InitiateDeposit(ctx context.Context, request model.InitiateDepositRequest) (model.Deposit, error)
	GetAllDeposits(ctx context.Context) (model.DepositBatchResponse, error)
	GetDepositID(ctx context.Context, id uuid.UUID) (model.Deposit, error)
	InternalFundsTransfer(ctx context.Context, request model.FundTransferRequest) (model.Deposit, error)
	IntraTransfer(ctx context.Context, request model.IntraTransferRequest) (model.IntraTransferResponse, error)

	// Transfer API
	InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.Transfer, error)
	GetExchangeRates(ctx context.Context, request model.GetExchangeRateRequest) (model.ExchangeRateDetails, error)

	// Withdrawal APIs
	InitiateWithdrawal(ctx context.Context, request model.InitiateWithdrawalRequest) (model.Withdrawal, error)
	FiatWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error)
	CryptoWithdrawal(ctx context.Context, request model.WithdrawalRequest) (model.Withdrawal, error)
	FeeWithdrawal(ctx context.Context, request model.FeeWithdrawalRequest) (model.FeeWithdrawal, error)

	// Wallet APIs
	GetWallet(ctx context.Context, request model.WalletRequest) (model.Wallet, error)
	GetWallets(ctx context.Context, customerID string) ([]*model.Wallet, error)
	GetSupportedAssets(ctx context.Context) ([]*model.SupportedAsset, error)

	// Transaction APIs
	GetTransactions(ctx context.Context, request *model.TransactionRequest) (model.TransactionResponse, error)

	// Payment APIs
	GetBanks(ctx context.Context) ([]model.BankCodeResponse, error)
	ResolveBankAccount(ctx context.Context, request model.AccountResolveRequest) (model.AccountDetailResponse, error)
	GenerateBankAccount(ctx context.Context, request model.BankAccountRequest) (model.BankAccountResponse, error)
	GetBankAccount(ctx context.Context, customerID uuid.UUID) (model.BankAccountResponse, error)

	// Payout APIs
	GetPayoutByID(ctx context.Context, payoutID string) (model.PayoutResponse, error)
	InitiateDirectBulkPayout(ctx context.Context, request model.InitiateBulkPayoutRequest) (model.PayoutDetails, error)
	InitiatePayout(ctx context.Context, currency, payoutType, beneficiaryType, remarks string, document *os.File) (model.PayoutDetails, error)
	GetAllPayouts(ctx context.Context, status, search string, dateBetween model.DateBetween, page model.Page) (model.AllPayoutsResponse, error)
	CancelPayout(ctx context.Context, request model.CancelPayoutRequest) error
	UpdatePayoutAccount(ctx context.Context, payoutID string, request model.TransferBeneficiaryDetails) error
	GetPayoutConfig(ctx context.Context, currency string) (model.BulkPayoutConfig, error)
	GetPayoutDocumentTemplate(ctx context.Context, currency, docType string) (string, error)

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
