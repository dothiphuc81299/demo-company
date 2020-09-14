package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// CompanyCreatePayload ...
	CompanyCreatePayload struct {
		Name            string  `json:"name"`
		CashbackPercent float64 `json:"cashbackPercent"`
	}
)

// Validate ...
func (payload CompanyCreatePayload) Validate() error {
	err := validation.Errors{
		"name":            validation.Validate(payload.Name, validation.Required, validation.Length(3, 20)),
		"cashbackPercent": validation.Validate(payload.CashbackPercent, validation.Required),
	}.Filter()
	return err
}
