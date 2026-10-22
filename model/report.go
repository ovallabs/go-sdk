package model

type (
	// SubmitSTRRequest schema for submitting a Suspicious Transaction Report
	SubmitSTRRequest struct {
		TransactionReference            string   `json:"transaction_reference"`
		SuspicionTypeCodes              []string `json:"suspicion_type_codes"`
		DescriptionOfSuspiciousActivity string   `json:"description_of_suspicious_activity"`
		ActionTaken                     string   `json:"action_taken"`
		PoliticallyExposedPerson        bool     `json:"politically_exposed_person"`
		TransactionCompleted            bool     `json:"transaction_completed"`
		ReasonNotCompleted              string   `json:"reason_not_completed"`
		// PublicPrivatePartnershipProjectCodes is empty when no FINTRAC PPP project applies.
		PublicPrivatePartnershipProjectCodes []string `json:"public_private_partnership_project_codes"`
	}
)
