package repository

import (
	"database/sql"

	"yahtzmen/core/port/repository"
)

type database struct {
	*sql.DB
}

func (d database) GetDB() *sql.DB {
	return d.DB
}
func (d database) Close() error {
	return d.DB.Close()
}

func NewDB() (repository.Database, error) {
	db, err := newDatabase()
	if err != nil {
		return nil, err
	}
	return &database{db}, nil
}

func newDatabase() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		"root:my-secret-pw@/yahtzmen",
	)
	if err != nil {
		return nil, err
	}
	return db, err
}
