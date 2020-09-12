package controllers

import (
	"github.com/labstack/echo/v4"

	"demo-company/config"
	"demo-company/modules/database"
	"demo-company/modules/zookeeper"
	"demo-company/validations"
)

// ServerHeader ....
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

// InitEcho ..
func InitEcho() *echo.Echo {
	e := echo.New()
	config.InitENV()
	envVars := config.GetEnv()
	zookeeper.Connect()
	database.Connect(envVars.Database.TestName)
	e.Use(ServerHeader)
	Boostrap(e)
	return e
}

// Boostrap ...
func Boostrap(e *echo.Echo) {
	e.POST("companies", CompanyCreate, validations.CompanyCreate)
	e.POST("branches",BranchCreate, validations.BranchCreate)
}

