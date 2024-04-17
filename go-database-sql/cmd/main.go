package main

import (
	db_service "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-database-sql/internal"
)

func main() {
	db_service.Conn()
	db_service.FetchData()
}
