// Package example holds sample models for requests
package example

import (
	"github.com/google/uuid"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

var (
	// NewCreateCustomerRequest sample create customer request
	NewCreateCustomerRequest = model.CreateCustomerRequest{
		Name:             "Nonso Adedayo",
		Email:            "chinonso@ovalfinanc.com",
		Reference:        "ref1230",
		MobileNumber:     "090803406089",
		Type:             model.CustomerTypeIndividual,
		YieldOfferingIDs: []uuid.UUID{uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77")},
	}

	// NewUpdateCustomerRequest sample update customer request
	NewUpdateCustomerRequest = model.UpdateCustomerRequest{
		CustomerID:       "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Name:             "Chinonso Okoli",
		Email:            "chinonso@ovalfinance.com",
		Reference:        "ref123",
		MobileNumber:     "09080340609",
		YieldOfferingIDs: []uuid.UUID{uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77")},
	}

	// NewCreateYieldOfferingProfilesRequest creates a CreateYieldOfferingProfilesRequest struct to use as an example
	NewCreateYieldOfferingProfilesRequest = model.CreateYieldOfferingProfilesRequest{
		Name:                  "aave yield",
		Description:           "for aave",
		APYRate:               10,
		Currency:              "usd",
		DepositLockDay:        2,
		MinimumDepositAllowed: 100,
		MaximumDepositAllowed: 3000,
		YieldableAfterDay:     1,
		WithdrawalLimitRate:   900,
		PortfolioID:           "c7115f87-11aa-4d69-bcb4-c12dd7f5bf2f",
		Reference:             "ref120",
	}

	// NewUpdateYieldOfferingProfilesRequest creates a UpdateYieldOfferingProfilesRequest struct to use as an example
	NewUpdateYieldOfferingProfilesRequest = model.UpdateYieldOfferingProfilesRequest{
		YieldOfferingID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
		Name:            "aave yield new name",
		Description:     "for aave new name oh... nothing again",
	}

	// NewGetYieldProfileByIDRequest creates a GetYieldProfileByIDRequest struct to use as an example
	NewGetYieldProfileByIDRequest = model.GetYieldProfileByIDRequest{
		YieldProfileID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
	}

	// NewDepositRequest newDepositRequest model
	NewDepositRequest = model.InitiateDepositRequest{
		CustomerID: "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Reference:  "ref123",
		Amount:     300,
	}

	// NewTransferRequest newTransferRequest model
	NewTransferRequest = model.InitiateTransferRequest{
		CustomerID: "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Amount:     20,
		Currency:   "USD",
		Destination: model.TransferDestination{
			BankDetails: model.BankDetails{
				AccountNumber: "11094843943",
				AccountName:   "Oval Banks",
				RoutingNumber: "3094395343",
				SwiftCode:     "",
				BankName:      "Oval US Investment Bank",
				BankBranch:    "",
				Country:       "US",
				City:          "",
				BankAddress:   "",
				District:      "",
				PostalCode:    "",
				IsWithinUS:    "yes",
			},
			PersonalDetails: model.PersonalDetails{
				Name:       "'Wale Oladapo",
				Country:    "GB",
				City:       "London",
				Address:    "London",
				District:   "",
				PostalCode: "304903",
			},
		},
		Note:      "",
		Reason:    "Gift token",
		Reference: "ref123",
	}

	//NewInitiateWithdrawalRequest newInitiateWithdrawalRequest model
	NewInitiateWithdrawalRequest = model.InitiateWithdrawalRequest{
		CustomerID: "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Reference:  "ref123",
		Amount:     210,
	}

	// NewInitiateBulkPayoutRequest sample bulk payout request
	NewInitiateBulkPayoutRequest = model.InitiateBulkPayoutRequest{
		Currency:        "NGN",
		Remarks:         "Some remarks",
		BeneficiaryType: model.SinglePayout,
		BeneficiaryID:   helpers.GetPointerString("57ef5467-5c19-4b1d-a8c1-5cb1f34bc587"),
		Amount:          helpers.GetPointerFloat64(1000),
	}

	// NewCancelPayoutRequest sample cancel payout request
	NewCancelPayoutRequest = model.CancelPayoutRequest{
		BulkPayoutID: "ef467f44-ed91-4875-8861-c2a5c7e4232d",
		Reason:       "Some reason",
	}

	// NewGenerateBankAccountRequest sample generate bank account request
	NewGenerateBankAccountRequest = model.GenerateBankAccountRequest{
		CustomerID:  "c4b9197f-009e-4019-b0dd-0cab6e9e3189",
		Currency:    "NGN",
		Reference:   "ref123",
		BVN:         helpers.GetPointerString("22000000000"),
		PhoneNumber: helpers.GetPointerString("2348109023376"),
	}

	// NewMockCustomerDepositRequest sample mock customer deposit request
	NewMockCustomerDepositRequest = model.MockCustomerDepositRequest{
		CustomerID: "c4b9197f-009e-4019-b0dd-0cab6e9e3189",
		Currency:   "NGN",
		Amount:     809000,
	}
)
