package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	Port   uint
	server *http.Server
}

func (h httpServer) Start() {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && !errors.Is(
			err,
			http.ErrServerClosed,
		) {
			log.Fatalf(
				"failed to stater HttpServer listen port %d, err=%s\n",
				h.Port,
				err.Error(),
			)
		}
	}()
	log.Printf(
		"Start Service with port %d\n",
		h.Port,
	)
}

func (h httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Fatalf(
			"Server forced to shutdown err=%s\n",
			err.Error(),
		)
	}
}

func NewHttpServer(router *echo.Echo) HttpServer {
	return &httpServer{
		Port: 8080,
		server: &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:8080"),
			Handler: router,
		},
	}
}
