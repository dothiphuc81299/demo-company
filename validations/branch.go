package validations

import (
	"github.com/labstack/echo/v4"

	"demo-company/models"
	"demo-company/dao"
	"demo-company/util"
)

// BranchCreate ..
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload   models.BranchCreatePayload
		)

		// Validate BranchCreatePayload
		c.Bind(&payload)
		err := payload.Validate()

		// If err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		companyID,err := util.HelperParseStringToObjectID(payload.Company)

		// if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		//check data
		company, _ := dao.CompanyFindByID(companyID)

		// if err
		if company.ID.IsZero() {
			return util.Response404(c, nil, err.Error())
		}

		payload.CompanyID = company.ID
		c.Set("branchPayload", payload)
		
		return next(c)
	}
}
