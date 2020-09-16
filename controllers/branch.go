package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-company/models"
	"demo-company/services"
	"demo-company/utils"
)

// BranchCreate ...
func BranchCreate(c echo.Context) error {
	var (
		payload = c.Get("branchPayload").(models.BranchCreatePayload)
	)

	// Process data
	rawData, err := services.BranchCreate(payload)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}
