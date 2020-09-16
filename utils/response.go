package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response ...
type Response map[string]interface{}

// generateResponse ...
func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

// Response200 ...
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "thanh cong!"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

// Response400 ...
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "du lieu khong hop le"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response404 ...
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = " du lieu khong tim thay"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}
