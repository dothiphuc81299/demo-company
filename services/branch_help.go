package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/models"
	"demo-company/util"
)

//  branchCreatePayloadToBSON ...
func branchCreatePayloadToBSON(payload models.BranchCreatePayload) models.BranchBSON {
	var (
		companyID, _ = util.ValidationObjectID(payload.CompanyID)
	)

	result := models.BranchBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return result

}
