package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/models"
	"demo-company/modules/database"
)

// BranchCreate ...
func BranchCreate(doc models.BranchBSON) (models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	// Insert branch
	_, err := branchCol.InsertOne(ctx, doc)
	return doc, err
}

// BranchFindByID ...
func BranchFindByID(id primitive.ObjectID) (models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
		result    models.BranchBSON
		filter    = bson.M{"_id": id}
	)

	// Find branch
	err := branchCol.FindOne(ctx, filter).Decode(&result)
	return result, err
}

// BranchUpdateByID ...
func BranchUpdateByID(filter bson.M, updateData bson.M) (err error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	// Update branch
	_, err = branchCol.UpdateOne(ctx, filter, updateData)
	return err
}
