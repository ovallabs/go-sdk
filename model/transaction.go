package model

type (
	// Transaction schema for transaction
	Transaction struct {
		ID              string      `json:"id"`
		BusinessID      string      `json:"businessID"`
		CustomerID      string      `json:"customerID"`
		YieldOfferingID string      `json:"yieldOfferingID"`
		Type            string      `json:"type"`
		Amount          float64     `json:"amount"`
		Currency        string      `json:"currency"`
		Reference       string      `json:"reference"`
		Status          string      `json:"status"`
		Destination     Destination `json:"destination"`
		CompletedAt     string      `json:"completedAt"`
		CreatedAt       string      `json:"createdAt"`
		BatchDate       string      `json:"batchDate"`
	}

	// AllTransactionsResponse schema for all transactions response
	AllTransactionsResponse struct {
		Items struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"items"`
		Page PageInfo `json:"page"`
	}
)
