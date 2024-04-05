package api

import (
	"database/sql"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"yahtzmen/internal/data"
)

type Application struct {
	Models data.Models
}

func main() {
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Credentials: credentials.NewCredentials(),
			},
		),
	)
	db, err := sql.Open(
		"mysql",
		"postgres://postgres:password@localhost/yahtzmen?sslmode=disable",
	)

	app := Application{
		Models: data.NewModels(
			db,
			sess,
		),
	}
}
