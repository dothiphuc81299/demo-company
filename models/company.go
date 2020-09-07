package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// CompanyBSON ...
	CompanyBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		Name             string             `bson:"name"`
		Address          string             `bson:"address"`
		CashbagPercent   float64            `bson:"cashbagPercent"`
		TotalTransaction float64            `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue"`
		CreatedAt        time.Time          `bson:"createdAt"`
	}

	//CompanyDetail ...
	CompanyDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		Address          string             `json:"address"`
		CashbagPercent   float64            `json:"cashbagPercent"`
		TotalTransaction float64            `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue"`
		CreatedAt        time.Time          `json:"createdAt"`
	}
)
