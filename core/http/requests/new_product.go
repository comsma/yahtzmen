package requests

import (
	"yahtzmen/core/domain"
)

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Dimensions  string `json:"dimensions"`
	Features    string `json:"features"`
	Notes       string `json:"notes"`
}

func (r *CreateProductRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Features:    r.Features,
		Notes:       r.Notes,
		Dimensions:  r.Dimensions,
	}
}
