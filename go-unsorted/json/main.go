package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	// UserId int `json:"userId"`
	UserId int `json:"-"`
	// UserId int
	// userId    int
	// ID int `json:"id"`
	ID int `json:"-"`
	// Title     string `json:"title"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	// Json decoding
	url := "https://jsonplaceholder.typicode.com/todos/2"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		/* 1.
		bodybytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := string(bodybytes)
		fmt.Println(data)
		*/

		// todoItem := Todo{}
		// todoItem := Todo{1, 1, "Release video on Go", false}
		todoItem := Todo{1, 1, "", false}

		// var b []byte

		/* 2.
		err = json.Unmarshal(bodybytes, &todoItem)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("Data from API: %+v\n", todoItem)
		*/

		/* 3.
		json.NewDecoder(resp.Body).Decode(&todoItem)
		fmt.Printf("Data from API: %+v\n", todoItem)
		*/

		/* 4.
		decoder := json.NewDecoder(resp.Body)
		// decoder.DisallowUnknownFields()
		if err = decoder.Decode(&todoItem); err != nil {
			log.Fatal("Disallow Unknown Fields: ", err)
			return
		}
		fmt.Printf("Data from API: %+v\n", todoItem)
		*/

		// Json encoding

		// todo, err := json.Marshal(todoItem)
		todo, err := json.MarshalIndent(todoItem, "", "\t")

		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(todo))
	}

}
