package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	db_service "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-database-sql/internal"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	DB_PASS := os.Getenv("DB_PASS")
	dbconn, err := sql.Open("mysql",
		"root:"+DB_PASS+"@tcp(127.0.0.1:3306)/demo")

	db_service.Conn(dbconn, err)
	db_service.FetchData(dbconn)
}
