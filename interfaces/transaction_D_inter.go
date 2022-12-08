package interfaces

import (
	"context"
	"transactions_mysql/model"
)

type TransactionDetailInter interface {
	FindTransaction_DByID(ctx context.Context, trx_id int) ([]model.Transactions_detail, error)
	InsertTransaction_D(ctx context.Context, transactionDetail model.Transactions_detail) (model.Transactions_detail, error)
}
