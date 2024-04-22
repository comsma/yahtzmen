package data

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	db            *sql.DB
	Products      ProductModel
	ProductImages ProductImageModel
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewModels(db *sql.DB) Models {
	return Models{
		db:            db,
		Products:      ProductModel{db: db},
		ProductImages: ProductImageModel{db: db},
	}
}

func NewModelsWithTx(db DBTX) Models {
	return Models{
		Products:      ProductModel{db: db},
		ProductImages: ProductImageModel{db},
	}
}
