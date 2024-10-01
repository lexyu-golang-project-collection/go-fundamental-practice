package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	fmt.Println("MYSQL")
}

func createClient() (*sql.DB, error) {
	panic("")
}
