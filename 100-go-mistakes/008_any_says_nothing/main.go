package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    int
	Message string
	Data    any
}

func handleResponse(code int, message string, data any) []byte {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil
	}
	return bytes
}

func main() {
	user := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  25,
	}
	userResp := handleResponse(200, "success", user)

	var result map[string]any
	if err := json.Unmarshal(userResp, &result); err != nil {
		panic(err)
	}

	if data, ok := result["Data"].(map[string]any); ok {
		name := data["Name"].(string)
		age := int(data["Age"].(float64))
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	products := []string{"phone", "laptop", "airpod"}
	productsResp := handleResponse(200, "success", products)

	var productsResult map[string]any
	if err := json.Unmarshal(productsResp, &productsResult); err != nil {
		panic(err)
	}

	if data, ok := productsResult["Data"].([]any); ok {
		for _, item := range data {
			product := item.(string)
			fmt.Println(product)
		}
	}
}
