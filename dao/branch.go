package dao

import(
	"context"
	
	"demo-company/models"
	"demo-company/modules/database"
)

// BranchCreate ...
func BranchCreate(doc models.BranchBSON) (models.BranchBSON, error) {
	var (
		branchCol =database.BranchCol()
		ctx =context.Background()
	)

	_, err :=branchCol.InsertOne(ctx,doc)

	return doc,err
}