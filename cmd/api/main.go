package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"yahtzmen/internal/data"
	"yahtzmen/internal/storage"
)

type Application struct {
	Models  data.Models
	Storage storage.Storage
}

func main() {
	db, err := sql.Open(
		"mysql",
		"root:password@tcp(localhost:3306)/yahtzmen",
	)
	if err != nil {
		fmt.Println(err)
	}

	strg := storage.NewS3Storage(
		"nyc3.digitaloceanspaces.com",
		"DO00Y3PVJHFAF2VNJTTD",
		"Mq6E9YAJmSGhO+ecQnmEt61UbLAqKQwV2AUDxfdnvaI",
	)

	app := Application{
		Models: data.NewModels(
			db,
		),
		Storage: strg,
	}

	e := echo.New()
	app.routes(e)

	fmt.Println(e.Start(":8088"))
}
