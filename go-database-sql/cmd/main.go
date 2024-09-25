package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	db_service "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-database-sql/internal"
)

func main() {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	DB_PASS := os.Getenv("DB_PASS")
	dbconn, err := sql.Open("mysql",
		"root:"+DB_PASS+"@tcp(127.0.0.1:3306)/demo")

	db_service.Conn(dbconn, err)
	defer dbconn.Close()
	db_service.Init(dbconn)

	/*
		var pk int64
		for i := 0; i < 5; i++ {
			price := rand.Float64()*(200-50) + 50

			isAvailable := rand.Intn(2) == 0

			// Create Product instance
			product := db_service.Product{
				Name:      "Book" + strconv.Itoa(i),
				Price:     price,
				Available: isAvailable,
			}
			pk = db_service.InsertProduct(dbconn, product)
		}
		fmt.Println("Primary Key:", pk)
	// */

	// productById := db_service.GetProductById(dbconn, 5555)
	// fmt.Printf("%+v", productById)

	products := db_service.GetProducts(dbconn)
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	// db_service.QueryDemo(dbconn)
	// db_service.PrepareDemo(dbconn)
	// db_service.SingleRowQuery(dbconn)
	// db_service.SingleRowPrepare(dbconn)
	// db_service.InsertData(dbconn)

}
