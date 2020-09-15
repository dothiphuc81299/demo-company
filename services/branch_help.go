package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/models"
	"demo-company/util"
)

func branchCreatePayloadToBSON(payload models.BranchCreatePayload) models.BranchBSON {
	var (
		companyID = util.HelperParseStringToObjectID(payload.CompanyID)
	)

	result := models.BranchBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return result
}
