package handler

import (
	"context"
	"database/sql"
	"strconv"
	"time"
	"transactions_mysql/interfaces"
	"transactions_mysql/model"
)

type TransactionHandler interface {
	GetAllProduct() ([]model.Products, error)
	GetAllVouchers() ([]model.Voucher, error)
	GetAllNumberTransaction() ([]model.Transaction, error)
	GetTransactionByNumber(trxNumber string) (model.Transaction, error)
	// IsValidProduct(data any) (model.Products, bool, error)
	IsValidProduct(name string, quantity int) (model.Products, bool, error)
	GenerateTransactionDetail(prodModel *model.Products, quantity int) model.Transactions_detail
	ValidateVoucher(code string) (model.Voucher, error)
	InputTransactionToDatabase(transaction *model.Transaction) error
	GenerateTransaction(transaction *model.Transaction, Voucher model.Voucher, CustomerName string, Email string, Phone string)
}

type transactionHandler struct {
	transactionInter       interfaces.TransactionInter
	transactionDetailInter interfaces.TransactionDetailInter
	productInter           interfaces.ProductInter
	voucherInter           interfaces.VoucherInter
}

func NewtransactionsHandler(trxHandler interfaces.TransactionInter, trxDhandler interfaces.TransactionDetailInter, produchHandler interfaces.ProductInter, voucherHandler interfaces.VoucherInter) *transactionHandler {
	return &transactionHandler{trxHandler, trxDhandler, produchHandler, voucherHandler}
}

func (trxH *transactionHandler) GetAllProduct() ([]model.Products, error) {
	var ctx = context.Background()
	list, err := trxH.productInter.FindAllProducts(ctx)

	if err != nil {
		// fmt.Println("Data Kosong")
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, err
	}

	return list, nil

}

func (trxH *transactionHandler) GetAllVouchers() ([]model.Voucher, error) {
	var ctx = context.Background()
	list, err := trxH.voucherInter.FindAllVoucher(ctx)

	if err != nil {
		// fmt.Println("Data Kosong")
		return nil, err
	}

	return list, nil

}

func (trxH *transactionHandler) GetAllNumberTransaction() ([]model.Transaction, error) {

	var ctx = context.Background()

	list, err := trxH.transactionInter.FindAllTransactionNumber(ctx)

	if err != nil {
		return nil, err
	}

	return list, nil

}

func (trxH *transactionHandler) GetTransactionByNumber(trxNumber string) (model.Transaction, error) {
	var ctx = context.Background()

	list, err := trxH.transactionInter.FindTransactionByNumber(ctx, trxNumber)

	if err != nil {
		return list, err
	}

	if err == sql.ErrNoRows {
		return list, err
	}

	listTransactionDetail, err := trxH.transactionDetailInter.FindTransaction_DByID(ctx, *list.GetTrxId())
	if err != nil {
		panic(listTransactionDetail)
	}

	list.SetTransactionD(listTransactionDetail)

	return list, nil
}

// Return name of product and error to validate the input of product

func (trxH *transactionHandler) IsValidProduct(name string, quantity int) (model.Products, bool, error) {
	var ctx = context.Background()
	var isValidQuantity bool

	var item model.Products
	var err error
	item, err = trxH.productInter.SearchProductByName(ctx, name)
	if quantity > 0 {
		isValidQuantity = true
	}

	return item, isValidQuantity, err
}

// generate product_name,price,quantity,total
func (trxH *transactionHandler) GenerateTransactionDetail(prodModel *model.Products, quantity int) model.Transactions_detail {
	var transaction_details model.Transactions_detail
	transaction_details.SetProductName(*prodModel.GetProductName())
	transaction_details.SetQuantity(quantity)
	transaction_details.SetPrice(*prodModel.GetProductValue())
	var price float32 = *transaction_details.GetPrice()
	var quantity_ int = *transaction_details.GetQuantity()
	transaction_details.SetTotal(price * float32(quantity_))
	transaction_details.SetProductId(*prodModel.GetProductId())

	return transaction_details
}

func (trxH *transactionHandler) ValidateVoucher(code string) (model.Voucher, error) {
	ctx := context.Background()
	var vouch model.Voucher

	vouch, err := trxH.voucherInter.FindVoucherByCode(ctx, code)
	// fmt.Println(vouch)

	if err == sql.ErrNoRows {
		return vouch, err
	}

	return vouch, err
}

// This function generate Transaction Number by its own, transaction number will always be different
func (trxH *transactionHandler) InputTransactionToDatabase(transaction *model.Transaction) error {

	ctx := context.Background()
	lastTransactionId, err := trxH.transactionInter.InsertTransaction(ctx, *transaction)
	if err != nil {
		return err
	}

	for _, v := range *transaction.GetTransactionD() {
		v.SetTransactionId(lastTransactionId)
		_, err := trxH.transactionDetailInter.InsertTransaction_D(ctx, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (trxH *transactionHandler) GenerateTransaction(transaction *model.Transaction, Voucher model.Voucher, CustomerName string, Email string, Phone string) {
	var totalQuantity int
	var totalPayment float32
	for _, v := range *transaction.GetTransactionD() {
		var quant_ int = *v.GetQuantity()
		totalQuantity += quant_
		var total_ float32 = *v.GetTotal()
		totalPayment += total_
	}
	// var date = fmt.Sprintln(time.Now().Format("2006-01-02"))
	// var date_string = fmt.Sprintln(time.Now())
	// var date, _ = time.Parse("2006-01-02", date_string)
	var date = time.Now()

	var disc float32 = *Voucher.GetVoucherValue()
	var totaldiscount float32
	var email_nullstring sql.NullString
	email_nullstring.Scan(Email)
	var phone_nullstring sql.NullString
	phone_nullstring.Scan(Phone)

	if disc > 0 && totalPayment > 300000 {
		// fmt.Println(disc)
		totaldiscount = totalPayment * disc
	} else {
		totaldiscount = 0
	}
	transaction.SetTransaction(generateTrxNumber(), CustomerName, email_nullstring, phone_nullstring, date, totalQuantity, totaldiscount, totalPayment, 0, *transaction.GetTransactionD())

	// fmt.Println(*transaction.GetDiscount())
}

func generateTrxNumber() string {
	var nomorTransaksi string = strconv.Itoa(time.Now().Nanosecond())

	return nomorTransaksi[:5]
}
