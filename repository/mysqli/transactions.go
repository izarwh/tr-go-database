package repository

import (
	"context"
	"database/sql"
	"fmt"
	"transactions_mysql/model"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (trxRepo *transactionRepository) FindAllTransactionNumber(ctx context.Context) ([]model.Transaction, error) {
	var sliceTransactions []model.Transaction

	// query := "select id,trx_number,customer_name,email,phone,quantity,discount,total,pay,date from transactions"
	query := "select id,trx_number,customer_name,email,phone,quantity,discount,total,pay,date from transactions"

	res, err := trxRepo.db.QueryContext(ctx, query)

	if err != nil {
		fmt.Println(err)
		return sliceTransactions, err
	}

	for res.Next() {
		var trans_model model.Transaction
		var trx_id, trx_number, cust_name, email, phone, date, quantity, discount, total, pay, _ = trans_model.GetTrx()
		// var voucher

		res.Scan(trx_id, trx_number, cust_name, email, phone, quantity, discount, total, pay, date)

		sliceTransactions = append(sliceTransactions, trans_model)
	}

	return sliceTransactions, err
}

func (trxRepo *transactionRepository) FindTransactionByNumber(ctx context.Context, trxNumber string) (model.Transaction, error) {
	query := "select id,trx_number,customer_name,email,phone,quantity,discount,total,pay,date from transactions where trx_number like ?"

	var transaction model.Transaction
	res, err := trxRepo.db.QueryContext(ctx, query, trxNumber)
	if err == sql.ErrNoRows {
		return transaction, err
	}
	for res.Next() {

		var id, trx_number, cust_name, email, phone, date, quantity, discount, total, pay, _ = transaction.GetTrx()

		res.Scan(id, trx_number, cust_name, email, phone, quantity, discount, total, pay, date)
	}
	return transaction, nil
}

func (trxRepo *transactionRepository) InsertTransaction(ctx context.Context, transaction model.Transaction) (int, error) {
	query := "insert into transactions (trx_number,customer_name,email,phone,quantity,discount,total,pay,date) values (?,?,nullif(?,''),nullif(?,''),?,?,?,?,?)"
	// query := "insert into transactions (trx_number,customer_name,email,phone,quantity,discount,total,pay,date) values (?,?,?,?,?,?,?,?,?)"

	var _, trx_number, cust_name, email, phone, date, quantity, discount, total, pay, _ = transaction.GetTrx()

	res, err := trxRepo.db.ExecContext(ctx, query, trx_number, cust_name, *email, *phone, quantity, discount, total, pay, date)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}
