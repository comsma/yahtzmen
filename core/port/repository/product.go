package repository

import (
	"context"

	"yahtzmen/core/domain"
)

type ProductRepository interface {
	GetProductList(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, id string) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product)
}
