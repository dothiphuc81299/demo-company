package routes

import (
	"github.com/labstack/echo/v4"

	"demo-company/controllers"
	"demo-company/validations"
)

// Company ...
func Company(e *echo.Echo) {
	routes := e.Group("companies")
	routes.POST("", controllers.CompanyCreate, validations.CompanyCreate)
	routes.GET("/:id/transactions", controllers.TransactionFindByCompanyID, validations.CompanyCheckExistedByCompanyID)
}
