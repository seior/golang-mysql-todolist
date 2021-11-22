package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func GetDatabase() *sql.DB {
	var db *sql.DB
	var err error

	if os.Getenv("env") == "test" {
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/"+"todolist_test")

		if err != nil {
			panic(err)
		}
	} else {
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/"+"todolist")
		if err != nil {
			panic(err)
		}
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db
}
