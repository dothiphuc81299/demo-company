package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// BranchCreatePayload ...
type BranchCreatePayload struct {
	CompanyID string `json:"companyID`
	Name      string `json:"name"`
}

// Validate ...
func (payload BranchCreatePayload) Validate() error {
	err := validation.Errors{
		"companyID": validation.Validate(payload.CompanyID, validation.Required, is.MongoID),
		"name":      validation.Validate(payload.Name, validation.Required, validation.Length(3, 20)),
	}.Filter()
	return err
}
