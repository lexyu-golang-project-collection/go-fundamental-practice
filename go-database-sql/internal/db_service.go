package db_service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id          int
	name        string
	dbconn, err = sql.Open("mysql",
		"root:P@ssw0rd@tcp(127.0.0.1:3306)/demo")
)

func Conn() {

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

func FetchData() {

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
	defer dbconn.Close()

}
