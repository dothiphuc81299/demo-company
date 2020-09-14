package dao

import (
	"context"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"demo-company/models"
	"demo-company/modules/database"
)

// CompanyCreate ...
func CompanyCreate(doc models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// InsertOne ...
	_, err := companyCol.InsertOne(ctx, doc)
	return doc, err
}

// CompanyFindByID ...
func CompanyFindByID(id primitive.ObjectID) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     models.CompanyBSON
		filter     = bson.M{"_id": id}
	)

	// Find
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	
	return result, err
}