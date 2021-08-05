package controller_echo

import (
	"net/http"
	apiecho "testproject/pkg/api_echo"
)

import (
	"github.com/labstack/echo/v4"
)

// @title testproject API Documentation
// @version 1.0

func AddCustomer(c echo.Context) error {

	result, _ := apiecho.AddCustomer(c)
	return c.JSON(http.StatusOK, result)
}

func GetCustomer(c echo.Context) error {

	result, _ := apiecho.GetCustomer(c)
	return c.JSON(http.StatusOK, result)
}
