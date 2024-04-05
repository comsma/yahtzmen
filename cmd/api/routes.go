package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// routes defines routes with handlers for application
func (a *Application) routes(e *echo.Router) error {
	e.Add(
		http.MethodGet,
		"/api/v1/products",
		a.productList,
	)

}
