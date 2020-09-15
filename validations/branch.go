package validations

import (
	"github.com/labstack/echo/v4"

	"demo-company/models"
	"demo-company/util"
)

// BranchCreate ..
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.BranchCreatePayload
		)

		// Validate BranchCreatePayload
		c.Bind(&doc)
		err := doc.Validate()

		// If err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		companyID := util.HelperParseStringToObjectID(doc.CompanyID)
		c.Set("companyID", companyID)
		c.Set("branchPayload", doc)
		return next(c)
	}
}
