package ports

import "yahtzmen/internal/application/domain"

type ProductImageRepoPort interface {
	GetForProduct(productId string) ([]*domain.ProductImage, error)
	Create(productId string) (*domain.ProductImage, error)
}
