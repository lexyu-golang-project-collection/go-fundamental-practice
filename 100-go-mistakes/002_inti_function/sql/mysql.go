package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	fmt.Println("MYSQL")
}

func CreateMySqlClient() (*sql.DB, error) {
	fmt.Println("Create MYSQL Client")

	return db, nil
}
