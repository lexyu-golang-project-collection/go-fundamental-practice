package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId int `json:"userId"`
	ID     int `json:"id"`
	// Title     string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		todoItem := Todo{}

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		if err = decoder.Decode(&todoItem); err != nil {
			log.Fatal("Disallow Unknown Fields: ", err)
			return
		}
		fmt.Printf("Data from API - NewDecoder() : %+v\n", todoItem)
	}
}
