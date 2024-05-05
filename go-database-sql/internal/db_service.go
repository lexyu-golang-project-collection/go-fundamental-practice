package db_service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

type Product struct {
	Nane      string
	Price     float64
	Available bool
}

func Conn(dbconn *sql.DB, err error) {

	if err != nil {
		log.Fatal(err)
	}

	err = dbconn.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ping OK")
	}
}

func QueryDemo(dbconn *sql.DB) {

	rows, err := dbconn.Query("select custid, name from customer where custid = ?", 100)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func PrepareDemo(dbconn *sql.DB) {
	stmt, err := dbconn.Prepare("select custid, name from customer where custid = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(103)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func SingleRowQuery(dbconn *sql.DB) {
	err := dbconn.QueryRow("select name from customer where custid = ?", 106).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func SingleRowPrepare(dbconn *sql.DB) {
	stmt, err := dbconn.Prepare("select name from customer where custid = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow(106).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func InsertData(dbconn *sql.DB) {
	stmt, err := dbconn.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Tester")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func Init(db *sql.DB) {
	err := createProductTable(db)
	if err != nil {
		log.Fatal(err)
	}
}

func createProductTable(db *sql.DB) error {

	query := `CREATE TABLE IF NOT EXISTS products(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		available BOOLEAN,
		create_at timestamp DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	return err
}

func insertProduct(db *sql.DB) {

}
