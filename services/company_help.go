package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/models"
)

// companyCreatePayload ...
func companyCreatePayloadToBSON(payload models.CompanyCreatePayload) models.CompanyBSON {
	result := models.CompanyBSON{
		ID:             primitive.NewObjectID(),
		Name:           payload.Name,
		CashbagPercent: payload.CashbagPercent,
		CreatedAt:      time.Now(),
	}
	return result
}