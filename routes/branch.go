package routes

import (
	"github.com/labstack/echo/v4"

	"demo-company/controllers"
	"demo-company/validations"
)

// Branch ...
func Branch(e *echo.Echo){
	routes :=e.Group("branches")
	routes.POST("",controllers.BranchCreate,validations.BranchCreate,companyCheckExistedByID)
}