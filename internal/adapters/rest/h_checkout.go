package rest

import (
	"github.com/labstack/echo/v4"
	"yahtzmen/internal/application/domain"
)

func (a *Adapter) CreateCheckoutSession(ctx echo.Context) error {
	var params CreateCheckoutSessionJSONRequestBody
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(400, err)
	}

	cart := domain.Cart{}

	for _, product := range *params.Products {
		cart.Products = append(cart.Products, &domain.CartItem{
			ProductId: product.Id,
			Quantity:  product.Quantity,
		})
	}

	secret, err := a.api.CreateCheckoutSession(cart)
	if err != nil {
		return ctx.JSON(500, err)
	}

	response := CheckoutSession{
		ClientSecret: secret,
	}

	return ctx.JSON(200, response)
}
