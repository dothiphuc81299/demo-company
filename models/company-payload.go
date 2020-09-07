package models

type (
	// CompanyCreatePayload ...
	CompanyCreatePayload struct {
		Name           string  `json:"name"`
		Address        string  `json:"address"`
		CashbagPercent float64 `json:"cashbagPercent"`
	}
)
