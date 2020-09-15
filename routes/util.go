package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-company/dao"
	"demo-company/util"
)

func companyCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			companyID = c.Get("companyID").(primitive.ObjectID)
		)

		company, _ := dao.CompanyFindByID(companyID)

		// Check existed
		if company.ID.IsZero() {
			return util.Response404(c, nil, "Khong tim thay company")
		}
		return next(c)
	}
}
