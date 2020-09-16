package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BranchCreatePayload ...
type BranchCreatePayload struct {
	CompanyID primitive.ObjectID `json:"companyID"`
	Company   string             `json:"company"`
	Name      string             `json:"name"`
}

// Validate ...
func (payload BranchCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Company, validation.Required.Error("company khong duoc trong"), is.MongoID.Error("companyID khong hop le")),
		validation.Field(&payload.Name, validation.Required.Error("ten khong duoc trong"), validation.Length(3, 20).Error("en phai co it nhat 3 ki tu ")),
	)
}

// ConvertToBSON ...
func (payload BranchCreatePayload) ConvertToBSON() BranchBSON {
	result := BranchBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: payload.CompanyID,
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return result
}
