package repository

import (
	"context"
	"database/sql"
	"transactions_mysql/model"
)

type transaction_DRepository struct {
	db *sql.DB
}

func NewTransaction_DRepository(db *sql.DB) *transaction_DRepository {
	return &transaction_DRepository{db}
}

func (trxDRepo *transaction_DRepository) FindTransaction_DByID(ctx context.Context, trx_id int) ([]model.Transactions_detail, error) {
	var sliceTransactionsDetail []model.Transactions_detail

	query := "select product_name,price,quantity,total from transaction_details where transaction_id = ?"

	res, err := trxDRepo.db.QueryContext(ctx, query, trx_id)

	if err != nil {
		return sliceTransactionsDetail, err
	}

	for res.Next() {
		var trans_Detail_model model.Transactions_detail
		var _, product_name, price, quantity, total = trans_Detail_model.GetTrxD()
		// var voucher

		res.Scan(product_name, price, quantity, total)

		sliceTransactionsDetail = append(sliceTransactionsDetail, trans_Detail_model)
	}

	return sliceTransactionsDetail, nil
}

func (trxDRepo *transaction_DRepository) InsertTransaction_D(ctx context.Context, transactionDetail model.Transactions_detail) (model.Transactions_detail, error) {
	query := "insert into transaction_details (transaction_id,product_id,product_name,price,quantity,total) values (?,?,?,?,?,?)"

	_, err := trxDRepo.db.ExecContext(ctx, query, transactionDetail.GetTransactionId(), transactionDetail.GetProductId(), transactionDetail.GetProductName(), transactionDetail.GetPrice(), transactionDetail.GetQuantity(), transactionDetail.GetTotal())

	if err != nil {
		return transactionDetail, err
	}

	if err != nil {
		return transactionDetail, err
	}

	return transactionDetail, nil
}
