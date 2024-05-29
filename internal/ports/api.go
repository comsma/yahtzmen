package ports

import (
	"mime/multipart"
	"yahtzmen/internal/application/domain"
)

type ApiPort interface {
	GetProduct(id string) (domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	AddProductImage(productId string, file *multipart.FileHeader) error
	CreateCheckoutSession(cart domain.Cart) (*string, error)
	AdminCreateProduct(product *domain.Product) error
	AdminDeleteProduct(id string) error
	AdminGetProduct(id string) (domain.Product, error)
	GetProductImages(productId string) ([]*domain.ProductImage, error)
}
