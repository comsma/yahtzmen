package repository

import (
	"context"
	"log"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"yahtzmen/core/domain"
	"yahtzmen/core/port/repository"
	"yahtzmen/model"
)

type productRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (p productRepository) GetProductList(ctx context.Context) ([]domain.Product, error) {
	var products []model.Product

	err := p.bun.NewSelect().Model(&products).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var dProducts []domain.Product

	for _, product := range products {
		dProducts = append(
			dProducts,
			domain.Product{
				Id:          product.Id,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Features:    product.Features,
				Notes:       product.Notes,
				Dimensions:  product.Dimensions,
			},
		)
	}

	return dProducts, nil
}

func (p productRepository) GetProduct(ctx context.Context, id string) (domain.Product, error) {
	product := new(model.Product)

	pk, err := strconv.ParseInt(
		id,
		10,
		64,
	)
	if err != nil {
		panic(err)
	}

	err = p.bun.NewSelect().Model(product).Where(
		"id = ?",
		pk,
	).Scan(ctx)
	if err != nil {
		log.Fatalf(
			"An error has occured: %v",
			err,
		)
	}
	return domain.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Features:    product.Features,
		Notes:       product.Notes,
		Dimensions:  product.Dimensions,
	}, nil
}

func (p productRepository) CreateProduct(ctx context.Context, product domain.Product) {
	newProduct := &model.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Notes:       product.Notes,
		Dimensions:  product.Dimensions,
		Features:    product.Features,
	}
	res, err := p.bun.NewInsert().Model(newProduct).Exec(ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func NewProductRepository(db repository.Database) repository.ProductRepository {
	bunDb := bun.NewDB(
		db.GetDB(),
		mysqldialect.New(),
	)
	return &productRepository{
		bun: bunDb,
		db:  db,
	}
}
