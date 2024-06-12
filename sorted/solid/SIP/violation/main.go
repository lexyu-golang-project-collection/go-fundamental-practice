package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save(db *sql.DB) {
	// Save user to the database
	result, err := db.Exec("INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName)
	if err != nil {
		log.Fatal(err)
	}

	pk, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Primary key:%d", pk)
	}
}

func conn(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ping OK")
	}
}

func createUserTable(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS users(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:"+"P@ssw0rd"+"@tcp(localhost:3306)/demo")

	return db, err
}

func main() {

	u := &User{
		FirstName: "DemoUser",
		LastName:  "For Test",
	}

	db, _ := connectToDatabase()
	defer db.Close()

	conn(db)

	createUserTable(db)

	u.Save(db)
}
