package templates

import (
	"fmt"
	"transactions_mysql/helper"

	tm "github.com/buger/goterm"
)

func (trxH *transactionTemplate) TampilVoucher() {
	helper.ClearScreen()
	fmt.Println(tm.MoveTo(tm.Background(tm.Color(tm.Bold("LIST VOUCHERS"), tm.WHITE), tm.MAGENTA), 10, 2))
	tm.Clear()
	width := 3 * 15

	vouchers, err := trxH.Transactionhandler.GetAllVouchers()
	if err != nil {
		panic(err)
	}

	heigh := len(vouchers) * 3
	box := tm.NewBox(width, heigh, 0)
	line := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(line, "ID\t | CodeTransaksi\t | Discount\n")
	for _, v := range vouchers {
		fmt.Fprintf(line, "%v\t | %v\t | %.0f%%\n", *v.GetVoucherId(), *v.GetVoucherCode(), *v.GetVoucherValue()*100)
	}
	fmt.Fprint(box, line)
	fmt.Println(box)
	tm.Clear()
}
