package apptest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demo-company/config"
	"demo-company/modules/database"
	"demo-company/modules/zookeeper"
	"demo-company/routes"
	"demo-company/util"
)

// InitServer ...
func InitServer() *echo.Echo {
	config.InitENV()
	zookeeper.Connect()
	database.Connect()
	util.HelperConnect()

	// New echo
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	// Route
	routes.Boostrap(e)

	return e
}
