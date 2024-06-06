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

	// NewInitiateTransferRequest sample transfer request
	NewInitiateTransferRequest = model.InitiateTransferRequest{
		CustomerID: "c4b9197f-009e-4019-b0dd-0cab6e9e3189",
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
		Reference: "12345678",
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

	// NewInitiateCurrencySwapRequest sample initiate currency swap request
	NewInitiateCurrencySwapRequest = model.InitiateCurrencySwapRequest{
		FromCurrency: "USD",
		ToCurrency:   "NGN",
		Amount:       1000,
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

	NewCreateBeneficiaryRequest = model.CreateBeneficiaryRequest{
		BankDetails: model.BankDetails{
			AccountNumber: "0762866445",
			AccountName:   "ADEDAYO OLAOLUWA OMOTOSO",
			BankName:      "Access Bank",
			BankCode:      "044",
			Country:       "Nigeria",
			IsWithinUS:    "no",
		},
		Currency: "NGN",
	}

	// NewInitiateDepositRequest sample initiate deposit request
	NewInitiateDepositRequest = model.InitiateDepositRequest{
		CustomerID:      "c4b9197f-009e-4019-b0dd-0cab6e9e3189",
		Reference:       "ref123",
		Amount:          100,
		YieldOfferingID: "63abda53-301f-44c3-bae1-447af643c593",
	}

	// NewFundTransferRequest sample fund transfer request
	NewFundTransferRequest = model.FundTransferRequest{
		CustomerID:      "c4b9197f-009e-4019-b0dd-0cab6e9e3189",
		Reference:       "ref123",
		Amount:          100000,
		Action:          model.Credit,
		YieldOfferingID: "4890133f-85f2-4b0d-8f26-b8707bc50b45",
	}

	// NewIntraTransferRequest sample intra transfer request
	NewIntraTransferRequest = model.IntraTransferRequest{
		Reference: "ref123",
		Amount:    100000,
		Sender: model.TransferParty{
			CustomerID:      "9f40fb69-64e3-4d23-853a-0243af155427",
			YieldOfferingID: "9f40fb69-64e3-4d23-853a-0243af155427",
		},
		Receiver: model.TransferParty{
			CustomerID:      "cba2fe07-f8bd-4c5a-97e5-923032f0467b",
			YieldOfferingID: "42ee80d8-2a95-419c-aad1-5643d306948e",
		},
	}
)
