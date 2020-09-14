package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	companies = "companies"
	branches  = "branches"
)

// CompanyCol ...
func CompanyCol() *mongo.Collection {
	return db.Collection(companies)
}

// BranchCol ...
func BranchCol() *mongo.Collection {
	return db.Collection(branches)
}
