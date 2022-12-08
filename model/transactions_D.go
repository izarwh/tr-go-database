package model

type Transactions_detail struct {
	transactionId int
	product_name  string
	productId     int
	price         float32
	quantity      int
	total         float32
}

func (t *Transactions_detail) GetTrxD() (*int, *string, *float32, *int, *float32) {
	return &t.productId, &t.product_name, &t.price, &t.quantity, &t.total
}

func (t *Transactions_detail) GetProductName() *string {
	return &t.product_name
}

func (t *Transactions_detail) GetPrice() *float32 {
	return &t.price
}

func (t *Transactions_detail) GetQuantity() *int {
	return &t.quantity
}

func (t *Transactions_detail) GetTotal() *float32 {
	return &t.total
}

func (t *Transactions_detail) GetTransactionId() *int {
	return &t.transactionId
}

func (t *Transactions_detail) GetProductId() *int {
	return &t.productId
}

func (t *Transactions_detail) SetProductName(name string) {
	t.product_name = name
}

func (t *Transactions_detail) SetPrice(price float32) {
	t.price = price
}

func (t *Transactions_detail) SetQuantity(quantity int) {
	t.quantity = quantity
}

func (t *Transactions_detail) SetTotal(total float32) {
	t.total = total
}

func (t *Transactions_detail) SetTransactionId(id int) {
	t.transactionId = id
}
func (t *Transactions_detail) SetProductId(id int) {
	t.productId = id
}
