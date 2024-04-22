package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Application) serverError(c echo.Context, err error) error {

	log.Println(err)
	return c.JSON(
		http.StatusInternalServerError,
		map[string]string{"error": "internal server error"},
	)
}
