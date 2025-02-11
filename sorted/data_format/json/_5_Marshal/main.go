package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	todoItem := Todo{1, 1, "Release video on Go", false}

	// Json encoding
	todo, err := json.Marshal(todoItem)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(todo))
}
