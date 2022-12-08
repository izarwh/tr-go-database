package main

import (
	"transactions_mysql/database"
	"transactions_mysql/templates"
)

func main() {
	db := database.GetConnection()
	defer db.Close()
	templates.MenuPenjualan(db)
}
