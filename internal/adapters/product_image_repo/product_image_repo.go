package product_image_repo

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jaevor/go-nanoid"
	"time"
	"yahtzmen/internal/application/domain"
)

type ProductImage struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"-"`
	ProductId    string    `json:"product_id"`
	DisplayOrder int32     `json:"display_order"`
}

type Adapter struct {
	db *sql.DB
}

func (a Adapter) generateId() (string, error) {
	gen, err := nanoid.Standard(20)
	if err != nil {
		return "", err

	}

	id := gen()

	return id, nil
}
func (a Adapter) Create(productId string) (*domain.ProductImage, error) {
	productImageEnt := ProductImage{
		ProductId: productId,
	}

	id, err := a.generateId()
	if err != nil {
		return nil, err
	}
	productImageEnt.Id = id

	err = a.insert(&productImageEnt)
	if err != nil {
		return nil, err
	}

	productImage := &domain.ProductImage{
		Id:    productImageEnt.Id,
		Order: productImageEnt.DisplayOrder,
	}

	return productImage, nil

}

func NewAdapter(db *sql.DB) *Adapter {
	return &Adapter{db: db}
}

func (a Adapter) GetForProduct(productId string) ([]*domain.ProductImage, error) {
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
	rows, err := a.db.QueryContext(
		ctx,
		query,
		productId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productImageEntities []*ProductImage
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
		productImageEntities = append(productImageEntities, productImage)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var productImages []*domain.ProductImage
	for _, productImageEntity := range productImageEntities {
		productImages = append(productImages, &domain.ProductImage{
			Id:    productImageEntity.Id,
			Order: productImageEntity.DisplayOrder,
		})
	}
	return productImages, nil
}

func (a Adapter) insert(productImage *ProductImage) error {
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
	_, err := a.db.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		switch {
		case err.(*mysql.MySQLError).Number == 1452:
			return domain.ErrEntityNotFound

		}
		return err
	}
	return nil
}
