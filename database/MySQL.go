package database

import (
	"database/sql"
	"log"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:qwer1234@tcp(139.155.30.83:3306)/izbasar?charset=utf8")
	if err != nil {
		log.Print("An error has been occurred when connecting to the database server:", err)
		return db
	} else {
		return db
	}
}
