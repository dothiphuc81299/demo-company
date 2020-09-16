package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"demo-company/config"
	"demo-company/models"
	"demo-company/modules/database"
)

var (
	companyIDString = "5f24d45125ea51bc57a8285c"
	companyID,_       = HelperParseStringToObjectID(companyIDString)
	company         = models.CompanyBSON{
		ID:               companyID,
		Name:             "PhucMars",
		CashbackPercent:  10.5,
		TotalRevenue:     10.3,
		TotalTransaction: 11,
	}
)

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperConnect ...
func HelperConnect() {
	var (
		envVars = config.GetEnv()
		client  = database.GetClient()
	)

	// Set Database for test ...
	db := client.Database(envVars.Database.TestName)
	fmt.Println("Database Connected to", envVars.Database.TestName)
	database.SetDB(db)
}

// HelperCompanyCreateFake ...
func HelperCompanyCreateFake() string {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// Insert company
	_, err := companyCol.InsertOne(ctx, company)
	if err != nil {
		log.Println(err)
	}
	return companyIDString
}
