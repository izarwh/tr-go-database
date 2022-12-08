package interfaces

import (
	"context"
	"transactions_mysql/model"
)

type TransactionInter interface {
	FindAllTransactionNumber(ctx context.Context) ([]model.Transaction, error)
	FindTransactionByNumber(ctx context.Context, trxNumber string) (model.Transaction, error)
	InsertTransaction(ctx context.Context, transaction model.Transaction) (int, error)
}
