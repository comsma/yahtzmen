package data

import (
	"database/sql"

	"yahtzmen/internal/storage"
)

type location string
type ProductImage struct {
	UUID         string `json:"uuid"`
	ProductId    int64  `json:"product_id"`
	DisplayOrder int    `json:"display_order"`
	File         File   `json:"file"`
}

type ProductImageModel struct {
	DB *sql.DB
	S3 storage.S3
}

func (m ProductImageModel) UploadImage(productImage ProductImage) error {
	m.Insert(productImage)
	productImage.File.Upload(m.S3.Uploader)
	return nil
}

func (m ProductImageModel) Insert(productImage ProductImage) error {
	// add a record to the database
}
