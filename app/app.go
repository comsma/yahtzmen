package app

import (
	"database/sql"
	"log"
	"yahtzmen/internal/data"
	"yahtzmen/internal/payment"
	"yahtzmen/internal/storage"
)

type App struct {
	Models  data.Models
	Storage storage.Storage
	Payment payment.Payment
}

func NewApp(config Config) App {

	db, err := sql.Open(
		"mysql",
		config.DSN,
	)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	return App{
		Models: data.NewModels(
			db,
		),
		Storage: storage.NewS3Storage(
			config.Storage.Endpoint,
			config.Storage.AccessKey,
			config.Storage.SecretKey,
		),
		Payment: *payment.NewPayment(),
	}

}
