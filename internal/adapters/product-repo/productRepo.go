package productRepo

import (
	"context"
	"database/sql"
	"time"
	"yahtzmen/internal/application/domain"
	"yahtzmen/internal/utils"
)

type Product struct {
	Id          string
	Name        string
	Price       float32
	Description string
	Notes       string
	Dimensions  string
	CreatedAt   time.Time
	Features    string
}

type Adapter struct {
	db *sql.DB
}

func NewAdapter(db *sql.DB) *Adapter {
	return &Adapter{db: db}
}

func (a Adapter) Get(id string) (domain.Product, error) {
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
	var productEntity Product
	err := a.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&productEntity.Id,
		&productEntity.Name,
		&productEntity.Price,
		&productEntity.Description,
		&productEntity.Notes,
		&productEntity.Dimensions,
		&productEntity.Features,
	)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		Price:       productEntity.Price,
		Description: productEntity.Description,
		Notes:       productEntity.Notes,
		Dimensions:  productEntity.Dimensions,
		CreatedAt:   time.Time{},
		Features:    productEntity.Features,
	}, nil
}

func (a Adapter) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	query := `
		SELECT id, name, price, description, notes, dimensions, features
		FROM products`
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	rows, err := a.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productEntities []*Product
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
		productEntities = append(productEntities, &product)
	}

	for _, productEntity := range productEntities {
		products = append(products, &domain.Product{
			Id:          productEntity.Id,
			Name:        productEntity.Name,
			Price:       productEntity.Price,
			Description: productEntity.Description,
			Notes:       productEntity.Notes,
			Dimensions:  productEntity.Dimensions,
			CreatedAt:   time.Time{},
			Features:    productEntity.Features,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (a Adapter) Create(product *domain.Product) error {
	id, err := utils.GenerateId()
	err = a.Insert(
		&Product{
			Id:          id,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Notes:       product.Notes,
			Dimensions:  product.Dimensions,
			Features:    product.Features,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (a Adapter) Insert(product *Product) error {
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
	_, err := a.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a Adapter) Delete(id string) error {

	query := `
		DELETE FROM products
		WHERE id = ?`
	args := []any{id}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	_, err := a.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return err
	}
	return nil
}
