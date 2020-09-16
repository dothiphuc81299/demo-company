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
		CashbackPercent  float64            `bson:"cashbackPercent"`
		TotalTransaction int64              `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue"`
		CreatedAt        time.Time          `bson:"createdAt"`
	}

	//CompanyDetail ...
	CompanyDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		CashbackPercent  float64            `json:"cashbackPercent"`
		TotalTransaction int64              `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue"`
		CreatedAt        time.Time          `json:"createdAt"`
	}

	// CompanyBrief ...
	CompanyBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
