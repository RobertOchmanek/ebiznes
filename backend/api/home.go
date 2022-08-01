package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {

	return c.JSON(http.StatusOK, "Home page of e-biznes backend server. Please go to https://ebiznesfrontendcontainer.azurewebsites.net to use the application.")
}
