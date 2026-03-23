package model

type (
	// SubmitSTRRequest schema for submitting a Suspicious Transaction Report
	SubmitSTRRequest struct {
		TransactionReference            string   `json:"transaction_reference"`
		SuspicionTypeCodes              []string `json:"suspicion_type_codes"`
		DescriptionOfSuspiciousActivity string   `json:"description_of_suspicious_activity"`
	}
)
