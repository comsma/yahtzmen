//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../../openapi.yaml
package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/url"
	"strings"
	"yahtzmen/internal/ports"
)

type Adapter struct {
	api  ports.ApiPort
	http *echo.Echo
}

func NewAdapter(api ports.ApiPort) *Adapter {

	return &Adapter{api: api}
}

func (a *Adapter) Run() {
	e := echo.New()
	a.http = e
	RegisterHandlersWithBaseURL(a.http, a, "/api/v1")
	url1, _ := url.Parse("http://localhost:5173/")

	a.http.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().URL.Path, "/api/v1") {
				return true
			}
			return false
		},
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: url1,
			},
		}),
	}))

	//a.http.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	//	Skipper:    nil,
	//	Root:       "frontend/dist",
	//	Index:      "index.html",
	//	HTML5:      true,
	//	Browse:     false,
	//	IgnoreBase: false,
	//	Filesystem: http.FS(web.Web),
	//}))
	a.http.Start(":8007")
}
