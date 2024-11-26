package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type (
	// CreateBeneficiaryRequest schema for create beneficiary request
	CreateBeneficiaryRequest struct {
		PersonalDetails  *PersonalDetails  `json:"personal_details,omitempty"`
		BankDetails      BankDetails       `json:"bank_details"`
		IntermediaryBank *IntermediaryBank `json:"intermediary_bank,omitempty"`
		Currency         string            `json:"destination_currency"`
		Nickname         *string           `json:"nickname,omitempty"`
		CustomerID       *string           `json:"customer_id,omitempty"`
	}

	// TransferBeneficiary schema for transfer beneficiary
	TransferBeneficiary struct {
		ID                  uuid.UUID      `json:"id"`
		BusinessID          uuid.UUID      `json:"business_id"`
		Name                string         `json:"name"`
		Reference           string         `json:"reference"`
		Details             postgres.Jsonb `json:"details"`
		DestinationCurrency string         `json:"currency"`
		ComplianceStatus    string         `json:"compliance_status"`
		Nickname            string         `json:"nickname"`
		CustomerID          *uuid.UUID     `json:"customer_id"`
		CreatedAt           time.Time      `json:"created_at"`
		UpdatedAt           *time.Time     `json:"updated_at"`
	}

	// AllBeneficiariesResponse schema for all beneficiaries response
	AllBeneficiariesResponse struct {
		Items *[]TransferBeneficiary `json:"items"`
		Page  PageInfo               `json:"page"`
	}
)
