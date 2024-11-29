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
	todoItem2 := Todo{2, 2, "Todo2", false}

	// Json encoding
	todo, err := json.MarshalIndent(todoItem, "", "\t")
	todo2, err := json.MarshalIndent(todoItem2, "", "\t")

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(todo))
	fmt.Println(string(todo2))

}
