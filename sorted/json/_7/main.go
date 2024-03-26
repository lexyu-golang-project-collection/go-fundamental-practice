package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	todoItem := Todo{1, 1, "", false}

	// Json encoding
	todo, err := json.MarshalIndent(todoItem, "", "\t")

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(todo))
}
