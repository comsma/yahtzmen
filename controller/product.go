package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"yahtzmen/core/domain"
	"yahtzmen/core/http/requests"
	"yahtzmen/core/http/resources"
	"yahtzmen/core/port/repository"
)

type productController struct {
	productRepo repository.ProductRepository
}

func (p productController) createProduct(c echo.Context) error {

	req := new(requests.CreateProductRequest)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	fmt.Println(req)
	err = c.Validate(req)
	if err != nil {
		return err
	}

	p.productRepo.CreateProduct(
		c.Request().Context(),
		req.ToDomain(),
	)

	return c.String(
		http.StatusOK,
		"ok",
	)
}
func (p productController) getProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := p.productRepo.GetProduct(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		resources.ProductResourceFromDomain([]domain.Product{product}),
	)
}

func (p productController) getProductList(c echo.Context) error {
	products, err := p.productRepo.GetProductList(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(
		http.StatusOK,
		resources.ProductResourceFromDomain(products),
	)
}

func NewProductController(e *echo.Echo, productRepo repository.ProductRepository) {

	controller := &productController{productRepo: productRepo}

	e.GET(
		"/api/product",
		controller.getProductList,
	)
	e.GET(
		"/api/product/:id",
		controller.getProduct,
	)
	e.POST(
		"/api/product",
		controller.createProduct,
	)
}
