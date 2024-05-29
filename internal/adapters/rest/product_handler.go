package rest

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"yahtzmen/internal/application/domain"
)

func (a *Adapter) GetProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) CreateProduct(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) ListStoreProducts(ctx echo.Context) error {

	products, err := a.api.GetAllProducts()

	if err != nil {
		return ctx.JSON(500, err)
	}

	var response []StoreProductPreview

	for _, product := range products {
		storeProductPreview := StoreProductPreview{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		if len(product.Images) > 0 {
			storeProductPreview.Image = ProductImage{Id: product.Images[0].Id, Url: product.Images[0].Url}
		}

		response = append(response, storeProductPreview)

	}
	return ctx.JSON(200, response)
}

func (a *Adapter) GetStoreProductById(ctx echo.Context, id string) error {
	product, err := a.api.GetProduct(id)

	if err != nil {
		return ctx.JSON(500, err)
	}

	response := StoreProduct{
		Description: &product.Description,
		Dimensions:  &product.Dimensions,
		Features:    &product.Features,
		Id:          &product.Id,
		Name:        &product.Name,
		Notes:       &product.Notes,
		Price:       &product.Price,
	}
	for _, image := range product.Images {
		response.Images = append(response.Images, ProductImage{Id: image.Id, Url: image.Url})
	}
	return ctx.JSON(200, response)

}

func (a *Adapter) UploadProductImage(ctx echo.Context, id string) error {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(400, "missing file")
	}
	err = a.api.AddProductImage(id, fileHeader)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEntityNotFound):
			return ctx.JSON(http.StatusBadRequest, "provided product id does not exist")
		}
		return ctx.JSON(500, err)

	}
	return ctx.JSON(200, "ok")
}
