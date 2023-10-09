package resources

import (
	"yahtzmen/core/domain"
)

type ProductResource struct {
	Data []*ProductResourceItem `json:"data"`
}

type ProductResourceItem struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Dimensions  string `json:"dimensions"`
	Features    string `json:"features"`
	Notes       string `json:"notes"`
}

func ProductResourceFromDomain(products []domain.Product) ProductResource {
	resource := new(ProductResource)

	for _, product := range products {
		resource.Data = append(
			resource.Data,
			&ProductResourceItem{
				Id:          product.Id,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Dimensions:  product.Dimensions,
				Features:    product.Features,
				Notes:       product.Notes,
			},
		)
	}

	return *resource
}
