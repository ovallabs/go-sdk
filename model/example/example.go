// Package example holds sample models for requests
package example

import (
	"github.com/google/uuid"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

var (
	// NewCreateCustomerRequest creates a CreateCustomerRequest struct to use as an example
	NewCreateCustomerRequest = model.CreateCustomerRequest{
		Name:             "Nonso Adedayo",
		Email:            "chinonso@ovalfinanc.com",
		Reference:        "ref1230",
		MobileNumber:     "090803406089",
		Type:             "individual",
		YieldOfferingIDs: []uuid.UUID{uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77")},
	}
	// NewUpdateCustomerRequest creates a UpdateCustomerRequest struct to use as an example
	NewUpdateCustomerRequest = model.UpdateCustomerRequest{
		CustomerID:       "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Name:             "Chinonso Okoli",
		Email:            "chinonso@ovalfinance.com",
		Reference:        "ref123",
		MobileNumber:     "09080340609",
		YieldOfferingIDs: []uuid.UUID{uuid.MustParse("ef8891af-e887-4e2c-ac79-7a9682d1ad77")},
	}

	// NewGetCustomerByIDRequest creates a GetCustomerByIDRequest struct to use as an example
	NewGetCustomerByIDRequest = model.GetCustomerByIDRequest{
		CustomerID: "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
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

	// NewInitiateTransferRequest sample transfer request
	NewInitiateTransferRequest = model.InitiateTransferRequest{
		CustomerID: "6cef5231-fc1e-45b3-a9ae-4d204245b0ae",
		Amount:     20000,
		Currency:   "NGN",
		Destination: model.TransferDestination{
			Type: "fiat",
			BankDetails: model.BankDetails{
				AccountNumber: "0762866445",
				AccountName:   "ADEDAYO OLAOLUWA OMOTOSO",
				BankName:      "Access Bank",
				Country:       "Nigeria",
				IsWithinUS:    "no",
			},
			PersonalDetails: model.PersonalDetails{
				Name:    "ADEDAYO OLAOLUWA OMOTOSO",
				Country: "Nigeria",
				City:    "Lagos",
				Address: "10 Balogun Street, Ikeja",
			},
		},
		Reason:    "Some reason",
		Reference: "ref124",
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

	NewInitiateTerminalTransferRequest = model.InitiateTerminalTransferRequest{
		Amount:              200,
		SourceCurrency:      "USD",
		DestinationCurrency: "NGN",
		UseBalance:          "yes",
		BeneficiaryID:       helpers.GetPointerString("c4158d8c-87a0-4f1b-b559-1aa2defd8495"),
		Note:                helpers.GetPointerString("Some note"),
		Reason:              "Some reason",
	}
)
