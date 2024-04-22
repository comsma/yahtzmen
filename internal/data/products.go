package data

import (
	"context"
	"time"

	"yahtzmen/internal/utils"
)

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Dimensions  string    `json:"dimensions"`
	CreatedAt   time.Time `json:"-"`
	Features    string    `json:"features"`
}

type ProductModel struct {
	db DBTX
}

func (m ProductModel) Create(name, description, notes, dimensions, features string, price float64) error {
	id, err := utils.GenerateId()
	err = m.Insert(
		&Product{
			Id:          id,
			Name:        name,
			Price:       price,
			Description: description,
			Notes:       notes,
			Dimensions:  dimensions,
			Features:    features,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (m ProductModel) Insert(product *Product) error {
	query := `
		INSERT INTO products (id, name, price, description, notes, dimensions, features)
		VALUES (?,?,?,?,?,?,?)`
	args := []any{
		product.Id,
		product.Name,
		product.Price,
		product.Description,
		product.Notes,
		product.Dimensions,
		product.Features,
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	_, err := m.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m ProductModel) Delete(product *Product) error {
	query := `
		DELETE FROM products
		WHERE id = ?`
	args := []any{product.Id}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	_, err := m.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m ProductModel) Get(id string) (*Product, error) {
	query := `
		SELECT id, name, price, description, notes, dimensions, features
		FROM products
		WHERE id = ?`
	args := []any{id}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	var product Product
	err := m.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Notes,
		&product.Dimensions,
		&product.Features,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m ProductModel) GetAll() ([]*Product, error) {
	query := `
		SELECT id, name, price, description, notes, dimensions, features
		FROM products`
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	rows, err := m.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Notes,
			&product.Dimensions,
			&product.Features,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
