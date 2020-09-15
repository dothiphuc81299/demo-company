package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-company/models"
	"demo-company/services"
	"demo-company/util"
)

// CompanyCreate ...
func CompanyCreate(c echo.Context) error {
	var (
		payload = c.Get("companyPayload").(models.CompanyCreatePayload)
	)

	// Process data
	rawData, err := services.CompanyCreate(payload)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

// TransactionFindByCompanyID ...
func TransactionFindByCompanyID(c echo.Context) error {
	var (
		companyID = c.Param("id")
	)

	// Process data
	rawData, err := services.TransactionFindByCompanyID(companyID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}
