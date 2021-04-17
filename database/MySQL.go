package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDB() *sql.DB {
	IpAddr := MYSQL_HOST
	port := MYSQL_PORT
	databaseName := MYSQL_DATABASE
	username := MYSQL_USERNAME
	password := MYSQL_PASSWORD
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, IpAddr, port, databaseName))
	if err != nil {
		log.Print("An error has been occurred when connecting to the database server:", err)
		return db
	} else {
		return db
	}
}
