package validations

import (
	"github.com/labstack/echo/v4"

	"demo-company/models"
	"demo-company/util"
)

// CompanyCreate ..
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.CompanyCreatePayload
		)

		// Validate CompanyCreatePayload
		c.Bind(&doc)
		err := doc.Validate()

		// If err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("companyPayload", doc)
		return next(c)
	}
}

// CompanyValidateID ...
func CompanyValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		// Validate companyID
		companyID, err := util.ValidationObjectID(id)

		// If err
		if err != nil {
			return util.Response400(c, nil, "ID khong hop le")
		}

		c.Set("companyID", companyID)
		return next(c)
	}
}
