package templates

import (
	"fmt"
	"transactions_mysql/helper"

	tm "github.com/buger/goterm"
)

func (trxH *transactionTemplate) TampilProduct() {
	helper.ClearScreen()
	fmt.Println(tm.MoveTo(tm.Background(tm.Color(tm.Bold("List Products"), tm.WHITE), tm.MAGENTA), 10, 2))
	tm.Clear()
	products, _ := trxH.Transactionhandler.GetAllProduct()
	// fmt.Println(products)
	// if err != sql.ErrNoRows {

	// }
	width := 60
	heigh := (len(products)+1)*2 - (len(products) - 1)

	box := tm.NewBox(width, heigh, 0)
	line := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(line, "ID\t | Item Name\t | Price\n")
	for _, v := range products {
		fmt.Fprintf(line, "%d\t | %s\t | %v\n", *v.GetProductId(), *v.GetProductName(), *v.GetProductValue())
	}
	fmt.Fprint(box, line)
	fmt.Println(box)
	tm.Clear()
}
