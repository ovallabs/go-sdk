// Package example holds sample models for requests
package example

import "github.com/ovalfi/go-sdk/model"

var (
	// NewCreateCustomerRequest creates a CreateCustomerRequest struct to use as an example
	NewCreateCustomerRequest = model.CreateCustomerRequest{
		Name:            "Nonso",
		Email:           "chinonso@ovalfinance.com",
		Reference:       "ref123",
		MobileNumber:    "09080340608",
		YieldOfferingID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
	}

	// NewUpdateCustomerRequest creates a UpdateCustomerRequest struct to use as an example
	NewUpdateCustomerRequest = model.UpdateCustomerRequest{
		CustomerID:      "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Name:            "Chinonso Okoli",
		Email:           "chinonso@ovalfinance.com",
		Reference:       "ref123",
		MobileNumber:    "09080340609",
		YieldOfferingID: "ef8891af-e887-4e2c-ac79-7a9682d1ad77",
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

	NewDepositRequest = model.InitiateDepositRequest{
		CustomerID: "cefec56e-3781-4b3a-bda6-ba4e7c0e49cd",
		Reference:  "ref123",
		Amount:     300,
	}

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
)
