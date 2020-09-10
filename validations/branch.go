package validations

import(
	"github.com/labstack/echo/v4"

	"demo-company/models"	
	"demo-company/util"
	"demo-company/services"

)

// BranchCreate ..
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.BranchCreatePayload
		)

		c.Bind(&doc)

		err :=doc.Validate()

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate object id
		companyID, err := services.StringToObjectID(doc.CompanyID)
		
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("branchPayload",doc)
		c.Set("companyID",companyID)

		return next(c)
	}
}