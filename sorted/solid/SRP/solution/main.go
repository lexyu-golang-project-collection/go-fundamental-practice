package main

import (
	"database/sql"
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

type UserRepository interface {
	Save(u *User)
}

type MySQLUserRepo struct {
	db *sql.DB
}

func (uRepo *MySQLUserRepo) Save(u *User) {
	stmt, _ := uRepo.db.Prepare("INSERT INTO users(first_name, last_name) VALUES(?,?)")

	res, _ := stmt.Exec(u.FirstName, u.LastName)
	lastId, _ := res.LastInsertId()
	rowCnt, _ := res.RowsAffected()

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func connectToDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:"+"P@ssw0rd"+"@tcp(localhost:3306)/demo")

	return db, err
}

func main() {

	u := &User{
		FirstName: "Tester",
		LastName:  "D",
	}
	db, _ := connectToDatabase()
	defer db.Close()

	repo := &MySQLUserRepo{db: db}

	repo.Save(u)
}
