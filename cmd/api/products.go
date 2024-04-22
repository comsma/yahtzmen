package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"yahtzmen/internal/data"
	"yahtzmen/internal/storage"
)

type ProductWithImage struct {
	data.Product
	PrimaryImage string `json:"primaryImage"`
}

func (a *Application) ProductListHandler(c echo.Context) error {
	products, err := a.Models.Products.GetAll()
	if err != nil {
		return a.serverError(c, err)
	}

	productsWithImages := make([]ProductWithImage, len(products))
	for i, product := range products {

		var id string
		primaryImage, err := a.Models.ProductImages.GetPrimaryForProduct(product.Id)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				id = "not-found"
			default:
				return a.serverError(c, err)
			}

		} else {
			id = primaryImage.Id
		}

		url, err := a.Storage.GetURL(c.Request().Context(), storage.ProductImagesPath, id)
		if err != nil {
			return a.serverError(c, err)
		}
		productsWithImages[i] = ProductWithImage{
			Product:      *product,
			PrimaryImage: url,
		}

	}
	return c.JSON(200, productsWithImages)
}

func (a *Application) ProductDeleteHandler(c echo.Context) error {
	id := c.Param("id")

	product, err := a.Models.Products.Get(id)
	if err != nil {
		return a.serverError(
			c,
			err,
		)

	}
	err = a.Models.Products.Delete(product)
	if err != nil {
		return a.serverError(
			c,
			err,
		)
	}
	return nil

}

func (a *Application) ProductImagePostHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return a.serverError(
			c,
			err,
		)
	}

	productId := c.Param("id")

	var order int32 = 1
	productImages, err := a.Models.ProductImages.GetAllForProduct(productId)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			order = 1
		default:
			return a.serverError(c, err)
		}
	}

	if len(productImages) > 0 {
		order = productImages[len(productImages)-1].DisplayOrder + 1
	}

	fmt.Println(order)

	err = a.Models.ExecInTx(
		func(m *data.Models) error {
			image, err := a.Models.ProductImages.Create(productId, order)
			if err != nil {
				return err
			}
			err = a.Storage.Upload(
				c.Request().Context(),
				storage.ProductImagesPath,
				image.Id,
				file,
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return a.serverError(
			c,
			err,
		)
	}

	return nil
}

func (a *Application) ProductCreateHandler(c echo.Context) error {
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Notes       string  `json:"notes"`
		Dimensions  string  `json:"dimensions"`
		Features    string  `json:"features"`
		Price       float64 `json:"price"`
	}
	if err := c.Bind(&input); err != nil {
		return a.serverError(
			c,
			err,
		)

	}
	err := a.Models.Products.Create(
		input.Name,
		input.Description,
		input.Notes,
		input.Dimensions,
		input.Features,
		input.Price,
	)
	if err != nil {
		return a.serverError(
			c,
			err,
		)
	}
	return nil
}

func (a *Application) ProductGetHandler(c echo.Context) error {

	type ProductImageWithUrl struct {
		data.ProductImage
		Url string `json:"url"`
	}
	id := c.Param("id")
	product, err := a.Models.Products.Get(id)
	productImages, err := a.Models.ProductImages.GetAllForProduct(id)
	if err != nil {
		return a.serverError(
			c,
			err,
		)
	}

	productImagesWithUrls := make([]ProductImageWithUrl, len(productImages))
	for i, productImage := range productImages {
		url, err := a.Storage.GetURL(c.Request().Context(), storage.ProductImagesPath, productImage.Id)
		if err != nil {
			return a.serverError(c, err)
		}
		productImagesWithUrls[i] = ProductImageWithUrl{
			ProductImage: *productImage,
			Url:          url,
		}
	}
	return c.JSON(
		200,
		map[string]interface{}{
			"product":       product,
			"productImages": productImagesWithUrls,
		},
	)
}
