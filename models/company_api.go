package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return validation.ValidateStruct(&payload,
		validation.Field(&payload.Name, validation.Required.Error("ten khong duoc trong"), validation.Length(3, 20).Error("ten phai co it nhat 3 ki tu ")),
		validation.Field(&payload.CashbackPercent, validation.Required.Error("cashbackPercent khong duoc trong")),
	)
}

// ConvertToBSON ...
func (payload CompanyCreatePayload) ConvertToBSON() CompanyBSON {
	result := CompanyBSON{
		ID:              primitive.NewObjectID(),
		Name:            payload.Name,
		CashbackPercent: payload.CashbackPercent,
		CreatedAt:       time.Now(),
	}
	return result
}
