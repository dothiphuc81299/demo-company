package validations

import (
	"github.com/labstack/echo/v4"

	"demo-company/dao"
	"demo-company/models"
	"demo-company/utils"
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
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("companyPayload", doc)
		return next(c)
	}
}

// CompanyCheckExistedByCompanyID ...
func CompanyCheckExistedByCompanyID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id = c.Param("id")
		)

		// Validate companyID
		companyID, err := utils.ValidationObjectID(id)

		// If err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		//check data
		company, err := dao.CompanyFindByID(companyID)

		// if err
		if company.ID.IsZero() {
			return utils.Response404(c, nil, err.Error())
		}
		
		c.Set("companyExisted", company)

		return next(c)
	}
}

