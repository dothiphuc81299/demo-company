package util

import (
	"bytes"
	"encoding/json"
	"io"
	"context"
	"log"
	"fmt"

	"demo-company/config"
	"demo-company/modules/database"
	"demo-company/models"
)

var (
	companyIDString ="5f24d45125ea51bc57a8285c"
	companyID = HelperParseStringToObjectID(companyIDString)
	company = models.CompanyBSON{
		ID:               companyID,
		Name:             "PhucMars",
		CashbackPercent:   10.5,
		TotalRevenue:     102.2,
		TotalTransaction: 16.2,
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

	// Insert
	_, err := companyCol.InsertOne(ctx,company)

	// err 
	if err != nil {
		log.Println(err)
	}

	return 	companyIDString
}