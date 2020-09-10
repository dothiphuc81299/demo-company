package services

import (
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/models"
)

// StringToObjectID ...
func StringToObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		fmt.Println(err)
	}
	return objectID, err
}

//  branchCreatePayloadToBSON ...
func branchCreatePayloadToBSON(payload models.BranchCreatePayload) models.BranchBSON{
	var (
		companyID,_ = StringToObjectID(payload.CompanyID)
	)

	result :=models.BranchBSON{
		ID:            primitive.NewObjectID(),
		CompanyID: companyID,
		Name: payload.Name,
		CreatedAt:time.Now(),
	}
	return result

}