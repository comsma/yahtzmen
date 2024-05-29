package api

import (
	"mime/multipart"
	"yahtzmen/internal/application/domain"
	"yahtzmen/internal/ports"
)

type Application struct {
	product      ports.ProductRepoPort
	productImage ports.ProductImageRepoPort
	file         ports.FilePort
	payment      ports.PaymentPort
}

func (a Application) GetProductImages(productId string) ([]*domain.ProductImage, error) {
	images, err := a.productImage.GetForProduct(productId)
	if err != nil {
		return nil, err
	}
	for _, image := range images {
		url, err := a.file.RetrieveFileUrl(image.Id)
		if err != nil {
			return nil, err
		}
		image.Url = url

	}
	return images, nil
}

func (a Application) AdminGetProduct(id string) (domain.Product, error) {
	product, err := a.product.Get(id)
	if err != nil {
		return domain.Product{}, err
	}
	product.Images, err = a.productImage.GetForProduct(product.Id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil

}
func (a Application) AdminDeleteProduct(id string) error {

	err := a.product.Delete(id)
	if err != nil {
		return err
	}
	return nil

}

func (a Application) AdminCreateProduct(product *domain.Product) error {
	err := a.product.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (a Application) CreateCheckoutSession(cart domain.Cart) (*string, error) {
	return a.payment.CreatePaymentIntent(100, "usd")
}

func (a Application) AddProductImage(productId string, file *multipart.FileHeader) error {

	productImage, err := a.productImage.Create(productId)
	if err != nil {
		return err
	}
	err = a.file.UploadFile(productImage.Id, file)
	if err != nil {

		return err
	}

	return nil
}

func (a Application) GetProduct(id string) (domain.Product, error) {
	product, err := a.product.Get(id)
	if err != nil {
		return domain.Product{}, err
	}
	product.Images, err = a.productImage.GetForProduct(product.Id)
	if err != nil {
		return domain.Product{}, err
	}
	for _, image := range product.Images {
		url, err := a.file.RetrieveFileUrl(image.Id)
		if err != nil {
			return domain.Product{}, err
		}
		image.Url = url
	}
	return product, nil
}

func (a Application) GetAllProducts() ([]*domain.Product, error) {
	products, err := a.product.GetAll()
	if err != nil {
		return nil, err

	}
	for _, product := range products {
		product.Images, err = a.productImage.GetForProduct(product.Id)
		if err != nil {
			return nil, err
		}
		for _, image := range product.Images {
			url, err := a.file.RetrieveFileUrl(image.Id)
			if err != nil {
				return nil, err
			}
			image.Url = url
		}

	}
	return products, nil
}

func NewApplication(product ports.ProductRepoPort, productImage ports.ProductImageRepoPort, file ports.FilePort, payment ports.PaymentPort) *Application {
	return &Application{product: product, productImage: productImage, file: file, payment: payment}
}
