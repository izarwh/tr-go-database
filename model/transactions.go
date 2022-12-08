package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	trx_id        int
	trx_number    string
	cust_name     string
	email         sql.NullString
	phone         sql.NullString
	date          time.Time
	quantity      int
	discount      float32
	total         float32
	pay           float32
	transaction_d []Transactions_detail //mempermudah saat memasukkan ke database
}

func (t *Transaction) GetTrx() (*int, *string, *string, *sql.NullString, *sql.NullString, *time.Time, *int, *float32, *float32, *float32, *[]Transactions_detail) {
	return &t.trx_id, &t.trx_number, &t.cust_name, &t.email, &t.phone, &t.date, &t.quantity, &t.discount, &t.total, &t.pay, &t.transaction_d
}

func (t *Transaction) GetTrxId() *int {
	return &t.trx_id
}

func (t *Transaction) GetTrxNumber() *string {
	return &t.trx_number
}

func (t *Transaction) GetCustName() *string {
	return &t.cust_name
}

func (t *Transaction) GetEmail() *sql.NullString {
	return &t.email
}

func (t *Transaction) GetPhone() *sql.NullString {
	return &t.phone
}

func (t *Transaction) GetDate() *time.Time {
	return &t.date
}

func (t *Transaction) GetQuantity() *int {
	return &t.quantity
}

func (t *Transaction) GetDiscount() *float32 {
	return &t.discount
}

func (t *Transaction) GetTotal() *float32 {
	return &t.total
}

func (t *Transaction) GetPay() *float32 {
	return &t.pay
}

func (t *Transaction) GetTransactionD() *[]Transactions_detail {
	return &t.transaction_d
}

func (t *Transaction) SetTrxId(trx_id int) {
	t.trx_id = trx_id
}

func (t *Transaction) SetTrxNumber(trx_number string) {
	t.trx_number = trx_number
}

func (t *Transaction) SetCustName(cust_name string) {
	t.cust_name = cust_name
}

func (t *Transaction) SetEmail(email sql.NullString) {
	t.email = email
}

func (t *Transaction) SetPhone(phone sql.NullString) {
	t.phone = phone
}

func (t *Transaction) SetDate(date string) {
	date_, _ := time.Parse("2006-01-02", date)
	t.date = date_
}

func (t *Transaction) SetQuantity(quantity int) {
	t.quantity = quantity
}

func (t *Transaction) SetDiscount(discount float32) {
	t.discount = discount
}

func (t *Transaction) SetTotal(total float32) {
	t.total = total
}

func (t *Transaction) SetPay(pay float32) {
	t.pay = pay
}

func (t *Transaction) SetTransactionD(transaction_d []Transactions_detail) {
	t.transaction_d = transaction_d
}

func (t *Transaction) SetTransaction(trx_number string, cust_name string, email sql.NullString, phone sql.NullString, date time.Time, quantity int, discount float32, total float32, pay float32, transaction_d []Transactions_detail) {
	t.trx_number = trx_number
	t.cust_name = cust_name
	t.email = email
	t.phone = phone
	t.date = date
	t.quantity = quantity
	t.discount = discount
	t.total = total
	t.pay = pay
	t.transaction_d = transaction_d
}
