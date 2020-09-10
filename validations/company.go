package validations

import(
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

		c.Bind(&doc)

		err :=doc.Validate()

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("companyPayload",doc)

		return next(c)
	}
}