package service

import (
	"yahtzmen/pkg/internal/core/model/response"
)

type ProductService interface {
	GetProductList() *response.GetProductListResponse
}
