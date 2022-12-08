package interfaces

import (
	"context"
	"transactions_mysql/model"
)

type ProductInter interface {
	FindAllProducts(ctx context.Context) ([]model.Products, error)
	SearchProductByName(ctx context.Context, data string) (model.Products, error)
}
