package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		todoItem := Todo{}
		err = json.Unmarshal(bodybytes, &todoItem)
		if err != nil {
			log.Print(err)
			return
		}

		fmt.Printf("Data from API: %+v\n", todoItem)
	}
}
