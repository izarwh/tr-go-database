package repository

import (
	"context"
	"database/sql"
	"transactions_mysql/model"
)

type prodRepository struct {
	db *sql.DB
}

func NewProdRepository(db *sql.DB) *prodRepository {
	return &prodRepository{db}
}

func (pRepo *prodRepository) FindAllProducts(ctx context.Context) ([]model.Products, error) {

	query := "select id, name, price from products"
	var sliceProduct []model.Products

	res, err := pRepo.db.QueryContext(ctx, query)

	if err != nil {
		return sliceProduct, err
	}

	for res.Next() {
		var product_model model.Products
		var id, name, value = product_model.GetProduct()

		res.Scan(id, name, value)
		sliceProduct = append(sliceProduct, product_model)
	}

	return sliceProduct, nil
}

func (trxRepo *prodRepository) SearchProductByName(ctx context.Context, data string) (model.Products, error) {
	query := "select id,name,price from products where name like ?"

	var foundProduct model.Products

	err := trxRepo.db.QueryRowContext(ctx, query, data).Scan(foundProduct.GetProductId(), foundProduct.GetProductName(), foundProduct.GetProductValue())

	if err == sql.ErrNoRows {
		return foundProduct, err
	}

	return foundProduct, err
}
