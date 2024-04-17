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

func FetchData(dbconn *sql.DB) {

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
