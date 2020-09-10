package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// BranchBSON ...
	BranchBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		CompanyID        primitive.ObjectID `bson:"companyID"`
		Name             string             `bson:"name"`
		TotalTransaction int64              `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue"`
		CreatedAt        time.Time          `bson:"createdAt"`
	}
)
