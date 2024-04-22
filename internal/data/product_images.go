package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"yahtzmen/internal/utils"
)

type location string
type ProductImage struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"-"`
	ProductId    string    `json:"product_id"`
	DisplayOrder int32     `json:"display_order"`
}

type ProductImageModel struct {
	db DBTX
}

func (m ProductImageModel) Create(productId string, displayOrder int32) (*ProductImage, error) {

	id, err := utils.GenerateId()
	if err != nil {
		return nil, err
	}

	productImage := &ProductImage{
		Id:           id,
		ProductId:    productId,
		DisplayOrder: displayOrder,
	}

	return productImage, m.Insert(productImage)
}

func (m ProductImageModel) GetPrimaryForProduct(productId string) (*ProductImage, error) {
	query := `
		SELECT id, product_id
		FROM product_images
		WHERE product_id = ?
		ORDER BY display_order ASC
		LIMIT 1`

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	row := m.db.QueryRowContext(
		ctx,
		query,
		productId,
	)

	productImage := &ProductImage{}
	err := row.Scan(
		&productImage.Id,
		&productImage.ProductId,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return productImage, nil

}

func (m ProductImageModel) GetAllForProduct(productId string) ([]*ProductImage, error) {
	query := `
		SELECT id, product_id, display_order
		FROM product_images
		WHERE product_id = ?
		ORDER BY display_order ASC`

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	rows, err := m.db.QueryContext(
		ctx,
		query,
		productId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productImages []*ProductImage
	for rows.Next() {
		productImage := &ProductImage{}
		err := rows.Scan(
			&productImage.Id,
			&productImage.ProductId,
			&productImage.DisplayOrder,
		)
		if err != nil {
			return nil, err
		}
		productImages = append(productImages, productImage)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return productImages, nil
}

func (m ProductImageModel) Insert(productImage *ProductImage) error {
	query := `
		INSERT INTO product_images (id, product_id, display_order)
		VALUES (?, ?, ?)`
	args := []any{
		productImage.Id,
		productImage.ProductId,
		productImage.DisplayOrder,
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
