// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"time"

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

	// Yield APIs
	GetBusinessPortfolios(ctx context.Context) ([]model.Portfolio, error)
	CreateYieldOfferingProfile(ctx context.Context, request model.CreateYieldOfferingProfilesRequest) (model.YieldOfferingProfile, error)
	UpdateYieldOfferingProfile(ctx context.Context, request model.UpdateYieldOfferingProfilesRequest) (model.UpdatedYieldOfferingProfile, error)
	GetAllYieldProfiles(ctx context.Context) ([]model.YieldOfferingProfile, error)
	GetYieldProfileByID(ctx context.Context, request model.GetYieldProfileByIDRequest) (model.YieldOfferingProfile, error)

	// Deposit APIs
	InitiateDeposit(ctx context.Context, request model.InitiateDepositRequest) (model.Deposit, error)
	GetAllDeposits(ctx context.Context) (model.DepositBatchResponse, error)
	GetDepositByBatchID(ctx context.Context, batchDate string) (model.Deposit, error)

	// Transfer API
	InitiateTransfer(ctx context.Context, request model.InitiateTransferRequest) (model.Transfer, error)

	// Withdrawal APIs
	InitiateWithdrawal(ctx context.Context, request model.InitiateWithdrawalRequest) (model.Withdrawal, error)

	// Wallet APIs
	GetWallet(ctx context.Context, request model.WalletRequest) (model.Wallet, error)
	GetWallets(ctx context.Context, customerID string) ([]*model.Wallet, error)
	GetSupportedAssets(ctx context.Context) ([]*model.SupportedAsset, error)

	// Run in sandbox mode
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
	c.SetTimeout(10 * time.Second)
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
