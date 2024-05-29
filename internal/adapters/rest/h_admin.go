package rest

import (
	"github.com/labstack/echo/v4"
	"yahtzmen/internal/application/domain"
)

func (a *Adapter) AdminGetProductImages(ctx echo.Context, id string) error {
	images, err := a.api.GetProductImages(id)
	if err != nil {
		return err
	}

	response := make([]ProductImage, len(images))
	for i, image := range images {
		response[i] = ProductImage{
			Id:    image.Id,
			Url:   image.Url,
			Order: int(image.Order),
		}

	}
	return ctx.JSON(200, response)
}

func (a *Adapter) AdminGetProductById(ctx echo.Context, id string) error {

	product, err := a.api.AdminGetProduct(id)
	if err != nil {
		return err
	}

	response := AdminProduct{
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		Description: product.Description,
		Dimensions:  product.Dimensions,
		Features:    product.Features,
		Id:          product.Id,
		IsLive:      product.IsLive,
		Name:        product.Name,
		Notes:       product.Notes,
		Price:       product.Price,
	}
	return ctx.JSON(200, response)

}

func (a *Adapter) AdminDeleteProduct(ctx echo.Context, id string) error {
	err := a.api.AdminDeleteProduct(id)
	if err != nil {
		return err

	}
	return ctx.JSON(200, nil)
}
func (a *Adapter) AdminGetProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) AdminCreateProduct(ctx echo.Context) error {
	var params NewAdminProduct
	if err := ctx.Bind(&params); err != nil {
		return err
	}

	product := &domain.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
		Notes:       params.Notes,
		Dimensions:  params.Dimensions,
		Features:    params.Features,
	}

	err := a.api.AdminCreateProduct(product)
	if err != nil {
		return err
	}
	return nil

}
