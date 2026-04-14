package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/comsma/yahtzmen/internal/config"
	"github.com/comsma/yahtzmen/internal/storage"
	"github.com/comsma/yahtzmen/internal/web"
	"github.com/comsma/yahtzmen/ui"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/urfave/cli/v3"
)

func serveCmd(cfg **config.Config) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Start the web server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: ":8080",
				Usage: "address to listen on",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return serve(*cfg, cmd.String("addr"))
		},
	}
}

func serve(cfg *config.Config, addr string) error {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()

	store, err := storage.New(
		cfg.Storage.Endpoint,
		cfg.Storage.Bucket,
		cfg.Storage.PublicBaseURL,
		cfg.Storage.AccessKey,
		cfg.Storage.SecretKey,
	)
	if err != nil {
		log.Fatalf("storage: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e := echo.New()
	renderer, err := web.NewTemplateCache()
	if err != nil {
		log.Fatalf("renderer: %v", err)
	}
	e.Renderer = renderer
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogMethod:   true,
		HandleError: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.Error("request", "method", v.Method, "uri", v.URI, "status", v.Status, "err", v.Error)
			} else {
				logger.Info("request", "method", v.Method, "uri", v.URI, "status", v.Status)
			}
			return nil
		},
	}))
	e.HTTPErrorHandler = func(c *echo.Context, err error) {
		if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
			if resp.Committed {
				return
			}
		}

		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		logger.Error("http error", "err", err, "status", code)

		if c.Request().Method == http.MethodHead {
			c.NoContent(code)
			return
		}

		tmpl := fmt.Sprintf("%d.gohtml", code)
		if rErr := c.Render(code, tmpl, nil); rErr != nil {
			// fallback: template not found for this code, use 500
			c.Render(http.StatusInternalServerError, "500.gohtml", nil)
		}
	}

	h := web.NewHandler(db, store)
	h.RegisterRoutes(e)

	e.GET("/static/*", func(c *echo.Context) error {
		return echo.StaticFileHandler("static/"+c.Param("*"), ui.Static)(c)
	})

	return e.Start(addr)
}
