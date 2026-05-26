package model

// CompetitorRate represents a normalized competitor exchange rate response.
type CompetitorRate struct {
	Provider     string  `json:"provider"`
	ProviderName string  `json:"provider_name"`
	LogoURL      string  `json:"logo_url"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Rate         float64 `json:"rate"`
}
