package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"yahtzmen/internal/adapters/file"
	"yahtzmen/internal/adapters/payment"
	productRepo "yahtzmen/internal/adapters/product-repo"
	"yahtzmen/internal/adapters/product_image_repo"
	"yahtzmen/internal/adapters/rest"
	"yahtzmen/internal/application/core/api"
)

func main() {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/yahtzmen")
	if err != nil {
		panic(err)
	}
	productRepoAdapter := productRepo.NewAdapter(db)
	productImageRepoAdapter := product_image_repo.NewAdapter(db)

	fileAdapter := file.NewAdapter("nyc3.digitaloceanspaces.com", "", "")

	paymentAdapter := payment.NewAdapter()

	app := api.NewApplication(productRepoAdapter, productImageRepoAdapter, fileAdapter, paymentAdapter)

	restServer := rest.NewAdapter(app)

	restServer.Run()

}
