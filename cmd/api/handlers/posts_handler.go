package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/theweird-kid/go-echo/cmd/api/services"
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

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to Process data!!")
	}
	data, err := services.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to Process data!!")
	}
	res := make(map[string]any)
	res["status"] = "OK"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}
