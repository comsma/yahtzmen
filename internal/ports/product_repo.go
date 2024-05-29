package ports

import "yahtzmen/internal/application/domain"

type ProductRepoPort interface {
	Get(id string) (domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Create(product *domain.Product) error
	Delete(id string) error
}
