package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	var driver_name = "mysql"
	var source_connection = "root:@tcp(localhost:3306)/transactions?parseTime=true"
	db, err := sql.Open(driver_name, source_connection)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(30 * time.Second)
	db.SetConnMaxLifetime(12 * time.Minute)
	return db
}
