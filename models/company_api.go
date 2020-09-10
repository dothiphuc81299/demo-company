package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// CompanyCreatePayload ...
	CompanyCreatePayload struct {
		Name           string  `json:"name"`
		CashbagPercent float64 `json:"cashbagPercent"`
	}
)

// Validate ...
func (payload CompanyCreatePayload) Validate() error {
	err := validation.Errors{
		"name":           validation.Validate(payload.Name, validation.Required, validation.Length(3, 20)),
		"cashbagPercent": validation.Validate(payload.CashbagPercent, validation.Required),
	}.Filter()
	return err
}
