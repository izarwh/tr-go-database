package templates

import (
	"database/sql"
	"fmt"
	"transactions_mysql/helper"
	"transactions_mysql/model"
)

func (t *transactionTemplate) AllTransaction(AllTrx []model.Transaction) {
	fmt.Println("List Transaksi")
	for i, v := range AllTrx {
		fmt.Printf("%d %s\n", i+1, *v.GetTrxNumber())
	}
}

func (t *transactionTemplate) TampilTransaksi() {
	helper.ClearScreen()
	list, err := t.Transactionhandler.GetAllNumberTransaction()

	if err != nil {
		panic(err)
	}

	if err == sql.ErrNoRows {
		fmt.Println("Tidak ada transaksi")
	} else {
		t.AllTransaction(list)
		fmt.Println("Masukkan nomor Transaksi: ")
		var nomorTransaksi string
		fmt.Scanln(&nomorTransaksi)
		fmt.Println("Nomor Transaksi: ", nomorTransaksi)
		transaction, err := t.Transactionhandler.GetTransactionByNumber(nomorTransaksi)
		if err == sql.ErrNoRows {
			fmt.Println("Data tidak ditemukan")
		}
		fmt.Println("Id Transaksi: ", *transaction.GetTrxId())
		date_ := transaction.GetDate().Format("2006-01-02")

		fmt.Printf("Tanggal Transaksi: %s\n", date_)
		var email_ = transaction.GetEmail()
		fmt.Println("Email :", email_.String)
		var phone_ = transaction.GetPhone()
		fmt.Println("Phone: ", phone_.String)
		for _, v := range *transaction.GetTransactionD() {
			fmt.Println(*v.GetProductName(), *v.GetQuantity(), *v.GetTotal())
		}

	}

}
