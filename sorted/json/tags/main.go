package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id" db:"id" form:"user_id" binding:"required"`
	Username string `json:"username" db:"username" form:"username" binding:"required"`
	Email    string `json:"email" db:"email" form:"email" binding:"required"`
}

func main() {
	user := User{
		ID:       1,
		Username: "example",
		Email:    "example@example.com",
	}

	user2 := User{
		ID:       1,
		Username: "example",
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))

	jsonData2, _ := json.Marshal(user2)
	fmt.Println(string(jsonData2))
}
