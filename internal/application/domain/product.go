package domain

import "time"

type Product struct {
	Id          string
	Name        string
	Price       float32
	Description string
	Notes       string
	Dimensions  string
	CreatedAt   time.Time
	Features    string
	IsLive      bool
	Images      []*ProductImage
}

type ProductImage struct {
	Id    string
	Url   string
	Order int32
}
