package templates

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"transactions_mysql/helper"
	"transactions_mysql/model"

	tm "github.com/buger/goterm"
)

//sesiTransaction
//input nama customer

func (trxH *transactionTemplate) AddTransaction() {
	//input data yang dibeli

	//input voucher
	//verifikasi apakah voucher ada > cari menggunakan findVoucherByCode
	//hitung total dari transaksi detail yang
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Tambah Penjualan/Transaksi")

	fmt.Println("Masukkan nama Customer: ")
	scanner.Scan()
	var namaCustomer = scanner.Text()
	fmt.Println("Nomor Telfon Customer: ")
	scanner.Scan()
	var nomorTelfon = scanner.Text()
	fmt.Println("Email Customer: ")
	scanner.Scan()
	var emailCustomer = scanner.Text()

	//store all validated data then store the transaction
	//first we need to loop all the input of our items
	//then we return the slice that we've already stored
	var trx model.Transaction
	// var sliceTransactionDetail []model.Transactions_detail
outer:
	for {
		trxH.TampilProduct()
		fmt.Println("Masukkan Nama Produk")
		scanner.Scan()
		var namaProduk = scanner.Text()
		fmt.Println("Masukkan Jumlah Produk")
		var jumlahProduk int
		fmt.Scanln(&jumlahProduk)

		// var prods_model, _, isValidName = trxH.Transactionhandler.IsValidProduct(&namaProduk)
		// var _, isValidQuantity, _ = trxH.Transactionhandler.IsValidProduct(&jumlahProduk)

		var prods_model, isValidQuantity, isValidName = trxH.Transactionhandler.IsValidProduct(namaProduk, jumlahProduk)

		//If input valid then complete all the struct of transactionDetail
		if (isValidName != sql.ErrNoRows) && (isValidQuantity) {
			//make struct of transactionDetail then we add slice of transactionDetail
			//for further transaction assignment
			trxDetail := trxH.Transactionhandler.GenerateTransactionDetail(&prods_model, jumlahProduk)
			*trx.GetTransactionD() = append(*trx.GetTransactionD(), trxDetail)
		} else {
			fmt.Println("Nama Produk atau Jumlah tidak Valid")
		}

		fmt.Println("Input data kembali? (y/n)")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "n":
			break outer
		case "y":
			continue
		default:
			break outer
		}
	}

	//if number of slice exceed the minimum, transaction will be created and submited to database
	if len(*trx.GetTransactionD()) > 0 {
		fmt.Println("Masukkan kode voucher")
		var voucher string
		fmt.Scanln(&voucher)
		var available_voucher, _ = trxH.Transactionhandler.ValidateVoucher(voucher)

		if *available_voucher.GetVoucherCode() == "" {
			fmt.Println("Voucher Invalid")
		}
		trxH.Transactionhandler.GenerateTransaction(&trx, available_voucher, namaCustomer, emailCustomer, nomorTelfon)

		fmt.Println("Total Belanja: ", *trx.GetTotal())
		fmt.Println("Total Diskon: ", *trx.GetDiscount())
		// time.Sleep(20 * time.Second)
		fmt.Println("Total Bayar: ", *trx.GetTotal()-*trx.GetDiscount())
		fmt.Println("Masukkan Jumlah Uang: ")
		var uang int
		fmt.Scanln(&uang)
		trx.SetPay(float32(uang))
		err := trxH.Transactionhandler.InputTransactionToDatabase(&trx)

		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Transaksi tidak valid")
	}
	showStruk(&trx)

}
func showStruk(trx *model.Transaction) {
	helper.ClearScreen()
	// width := reflect.TypeOf(model.Products[0]).NumField() * 10
	heigh := (len(*trx.GetTransactionD()) + 3) * 4
	box := tm.NewBox(60, heigh, 0)
	line := tm.NewTable(0, 10, 5, ' ', 0)

	// fmt.Fprintf(line, "%s\n", tm.MoveTo(teks, 15, 0))
	teks2 := tm.Background(tm.Color(tm.Bold("Struk Transaksi"), tm.WHITE), tm.MAGENTA)
	fmt.Fprintf(line, "%s\n", tm.MoveTo(teks2, 24, 2))

	fmt.Fprint(line, "Qty\t Item\t Total\n")
	for _, v := range *trx.GetTransactionD() {
		fmt.Fprintf(line, "%d\t %s\t %.0f\n", *v.GetQuantity(), *v.GetProductName(), *v.GetTotal())
	}
	fmt.Fprintf(line, "Total\t \t %v\n", *trx.GetTotal())
	fmt.Fprintf(line, "Discount\t %v\n", *trx.GetDiscount())
	var total_disc = *trx.GetTotal() - *trx.GetDiscount()
	fmt.Fprintf(line, "Total Discount:\t %v\n", total_disc)
	fmt.Fprintf(line, "Payment:\t %v\n", *trx.GetPay())
	fmt.Fprintf(line, "Change:\t %v\n", *trx.GetPay()-total_disc)
	date_ := trx.GetDate().Format("2006-01-02")
	fmt.Fprintf(line, "Transaction Date\t %s\n", date_)

	fmt.Fprint(box, line)
	fmt.Println(box)
}
