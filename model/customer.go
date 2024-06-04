package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	// CustomerTypeIndividual represent individual customer type
	CustomerTypeIndividual CustomerType = "individual"
	// CustomerTypeBusiness represent business customer type
	CustomerTypeBusiness CustomerType = "business"
)

type (
	// CustomerType customerType string
	CustomerType string

	// Customer schema for customer
	Customer struct {
		ID               string      `json:"id"`
		Name             string      `json:"customer_name"`
		MobileNumber     string      `json:"mobile_number"`
		Email            string      `json:"email"`
		Channel          string      `json:"channel"`
		Reference        string      `json:"reference"`
		YieldOfferingIDs []uuid.UUID `json:"api_yield_offering_ids"`
		UpdatedAt        *time.Time  `json:"updated_at"`
		CreatedAt        string      `json:"created_at"`
	}

	// CreateCustomerRequest schema for create customer request
	CreateCustomerRequest struct {
		Name             string       `json:"name"`
		Email            string       `json:"email"`
		Reference        string       `json:"reference"`
		MobileNumber     string       `json:"mobile_number"`
		Type             CustomerType `json:"type"`
		YieldOfferingIDs []uuid.UUID  `json:"yield_offering_ids"`
	}

	// UpdateCustomerRequest schema for update customer request
	UpdateCustomerRequest struct {
		CustomerID       string      `json:"customer_id"`
		Name             string      `json:"name"`
		Email            string      `json:"email"`
		Reference        string      `json:"reference"`
		MobileNumber     string      `json:"mobile_number"`
		YieldOfferingIDs []uuid.UUID `json:"yield_offering_ids"`
	}

	// AllCustomersResponse schema for all customers response
	AllCustomersResponse struct {
		Items []Customer `json:"items"`
		Page  PageInfo   `json:"page"`
	}

	// CustomerBalance schema for customer balance
	CustomerBalance struct {
		YieldOfferingID uuid.UUID `json:"yield_offering_id"`
		Name            string    `json:"name"`
		Currency        string    `json:"currency"`
		Amount          float64   `json:"balance"`
	}

	// CustomerBalances schema for customer balances
	CustomerBalances struct {
		CustomerID   uuid.UUID          `json:"customer_id"`
		TotalBalance float64            `json:"total_balance"`
		Detail       []*CustomerBalance `json:"detail"`
	}
)
