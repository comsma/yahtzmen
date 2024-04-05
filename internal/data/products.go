package data

import (
	"context"
	"database/sql"
	"time"
)

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Features    []string  `json:"features"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m ProductModel) CreateProduct(product Product) error {
	query := `
		INSERT INTO products (id, name, description)
		VALUES ($1, $2, $3)`
	args := []any{
		product.Id,
		product.Name,
		product.Description,
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	_, err := m.DB.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return err
	}
	return err
}
