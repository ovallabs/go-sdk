// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"github.com/ovalfi/go-sdk/model"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// RemoteCalls abstracted definition of supported functions
type RemoteCalls interface {
	CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (model.Customer, error)
	UpdateCustomer(ctx context.Context, request model.UpdateCustomerRequest) (model.Customer, error)
	GetAllCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomerByID(ctx context.Context, request model.GetCustomerByIDRequest) (model.CustomerInfo, error)
	RunInSandboxMode()
}

// Call object
type Call struct {
	client       *resty.Client
	logger       zerolog.Logger
	baseURL      string
	signature    string
	bearerToken  string
	sandboxMode  bool
	idempotentID uuid.UUID
}

// New initialises the object Call
func New(z *zerolog.Logger, c *resty.Client, signature, bearerToken, bURL string) RemoteCalls {
	c.SetTimeout(10 * time.Second)
	call := &Call{
		client:       c,
		logger:       z.With().Str("sdk", "ovalfi").Logger(),
		baseURL:      bURL,
		signature:    signature,
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
	//c.client.EnableTrace()
	c.sandboxMode = true
}
