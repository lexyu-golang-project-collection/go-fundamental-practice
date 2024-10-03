package main

import (
	"fmt"

	nosql "github.com/lexyu-golang-project-collection/go-fundamental-practice/100-go-mistakes/002_init_function/nosql"
	db "github.com/lexyu-golang-project-collection/go-fundamental-practice/100-go-mistakes/002_init_function/sql"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("Start...")

	k, v := nosql.Store("K", "V")
	fmt.Println(k, v)

	db.CreateMySqlClient()

}
