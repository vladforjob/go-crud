package config

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

//	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
/*	db, err := sql.Open("mysql", "root:root@/go_products_master?parseTime=true")
	if err != nil {
		panic(err)
	}

*/
	db, err := sql.Open("sqlite3", "./native.db")
	if err != nil {
		fmt.Println(err)
	}

	DB = db

	log.Println("Database connected")
}
