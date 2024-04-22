package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// routes defines routes with handlers for application
func (a *Application) routes(e *echo.Echo) error {
	fe, _ := url.Parse("http://localhost:5173/")
	e.Use(
		echomiddleware.ProxyWithConfig(
			echomiddleware.ProxyConfig{
				Skipper: func(c echo.Context) bool {
					if strings.HasPrefix(c.Request().URL.Path, "/api") {
						return true
					}
					return false
				}, Balancer: echomiddleware.NewRoundRobinBalancer(
					[]*echomiddleware.ProxyTarget{
						{
							URL: fe,
						},
					},
				),
			},
		),
	)
	e.Add(http.MethodGet, "/api/v1/products", a.ProductListHandler)
	e.Add(http.MethodGet, "/api/v1/products/:id", a.ProductGetHandler)
	e.Add(http.MethodPost, "/api/v1/products", a.ProductCreateHandler)
	e.Add(http.MethodDelete, "/api/v1/products/:id", a.ProductDeleteHandler)
	e.Add(http.MethodPost, "/api/v1/products/:id/image", a.ProductImagePostHandler)

	return nil

}
