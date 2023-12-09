package handlers

import (
	"net/http"

	"../services"
	"github.com/labstack/echo"
)

func PostIndexHandler(c echo.Context) error {
	data, err := services.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to Process data!!")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}
