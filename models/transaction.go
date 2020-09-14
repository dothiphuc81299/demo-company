package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionDetail ...
	TransactionDetail struct {
		ID                     primitive.ObjectID `json:"_id"`
		CompanyID              primitive.ObjectID `json:"companyID"`
		BranchID               primitive.ObjectID `json:"branchID"`
		UserID                 primitive.ObjectID `json:"userID"`
		Amount                 float64            `json:"amount"`
		Commission             float64            `json:"commission"`
		CompanyCashbackPercent float64            `json:"companyCashbackPercent"`
		CreatedAt              time.Time          `json:"createdAt"`
	}
)
