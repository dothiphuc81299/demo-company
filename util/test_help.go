package util

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"demo-company/modules/database"
	"demo-company/models"
)

var (
	company = models.CompanyBSON{
		ID:               primitive.NewObjectID(),
		Name:             "PhucMars",
		CashbackPercent:   10.5,
		TotalRevenue:     102.2,
		TotalTransaction: 16.2,
		CreatedAt:        time.Now(),
	}
	CompanyStringSuccess = "5f24d45125ea51bc57a8285c"
	CompanyStringInvalid = "5f24d45159"

)

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperCompanyCreateFake ...
func HelperCompanyCreateFake()  {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// Insert
	_, err := companyCol.InsertOne(ctx,company)

	if err != nil {
		log.Println("Error",err)
	}
}
