package data

import (
	"database/sql"

	"yahtzmen/internal/storage"
)

type Models struct {
	Products      ProductModel
	ProductImages ProductImageModel
}

func NewModels(db *sql.DB, s3 *storage.S3) Models {
	return Models{
		Products:      ProductModel{},
		ProductImages: ProductImageModel{},
	}
}
