package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	// CustomerWalletRequest request payload
	CustomerWalletRequest struct {
		CustomerID string `json:"customer_id"`
		Network    string `json:"network"`
		Asset      string `json:"asset"`
	}

	// CustomerWallet schema represents entity that contains customer wallet and other needed information
	CustomerWallet struct {
		ID            uuid.UUID  `json:"id"`
		CustomerID    uuid.UUID  `json:"customer_id"`
		BusinessID    uuid.UUID  `json:"business_id"`
		WalletAddress string     `json:"wallet_address"`
		Asset         string     `json:"asset"`
		Network       string     `json:"network"`
		Type          string     `json:"type"`
		Provider      string     `json:"provider"`
		Reason        string     `json:"reason"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedTime   *time.Time `json:"updated_at"`
	}
)
