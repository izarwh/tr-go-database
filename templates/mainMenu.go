package templates

import (
	"database/sql"
	"fmt"
	"transactions_mysql/handler"
	"transactions_mysql/helper"
	repository "transactions_mysql/repository/mysqli"

	tm "github.com/buger/goterm"
)

type transactionTemplate struct {
	Transactionhandler handler.TransactionHandler
	db                 *sql.DB
}

func NewTransactionTemplate(trxH handler.TransactionHandler, db *sql.DB) *transactionTemplate {
	return &transactionTemplate{trxH, db}
}

func MenuPenjualan(db *sql.DB) {

	//dependencies injecion
	// helper.ClearScreen()
	transactionsRepository := repository.NewTransactionRepository(db)
	transactionsDetailRepository := repository.NewTransaction_DRepository(db)
	productRepository := repository.NewProdRepository(db)
	voucherRepository := repository.NewVoucherRepository(db)

	transactionsHandler := handler.NewtransactionsHandler(transactionsRepository, transactionsDetailRepository, productRepository, voucherRepository)

	transactionTemplate := NewTransactionTemplate(transactionsHandler, db)

	teks := tm.Background(tm.Color(tm.Bold("Menu Transaksi"), tm.WHITE), tm.MAGENTA)
	fmt.Println(tm.MoveTo(teks, 20, 2))
	box := tm.NewBox(50, 8, 0)
	fmt.Fprintln(box, tm.Bold("1. Tambah Penjualan/Transaksi"))
	fmt.Fprintln(box, tm.Bold("2. Lihat Penjualan/Transaksi"))
	fmt.Fprintln(box, tm.Bold("3. Lihat list Product"))
	fmt.Fprintln(box, tm.Bold("4. Lihat list Voucher"))
	fmt.Fprintln(box, tm.Bold("0. Ke Luar"))

	fmt.Println(box)

	fmt.Println("Pilih menu: ")
	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		transactionTemplate.AddTransaction()
		helper.Backhandling()
		MenuPenjualan(db)
	case 2:
		transactionTemplate.TampilTransaksi()
		helper.Backhandling()
		MenuPenjualan(db)
	case 3:
		transactionTemplate.TampilProduct()
		helper.Backhandling()
		MenuPenjualan(db)
	case 4:
		transactionTemplate.TampilVoucher()
		helper.Backhandling()
		MenuPenjualan(db)
	// case 0:
	// 	os.Exit(3)
	default:
		MenuPenjualan(db)
	}

	//Buat struk
}
