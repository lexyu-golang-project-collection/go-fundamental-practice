package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Commonly used test API endpoints
const (
	JSONPLACEHOLDER_API = "https://jsonplaceholder.typicode.com"
	HTTPBIN_API         = "https://httpbin.org"
	DOG_API             = "https://dog.ceo/api"
	POKEMON_API         = "https://pokeapi.co/api/v2"
)

// Example data structures
type Post struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	Name     string `json:"name"`
	Job      string `json:"job"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GET Examples
func getExamples(client *http.Client) {
	// JSONPlaceholder - Get posts
	fmt.Println("\n--- JSONPlaceholder GET Example ---")
	resp, err := client.Get(JSONPLACEHOLDER_API + "/posts/1")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Posts Response: %s\n", string(body))
	}

	// Dog API - Get random dog image
	fmt.Println("\n--- Dog API GET Example ---")
	resp, err = client.Get(DOG_API + "/breeds/image/random")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Dog Image Response: %s\n", string(body))
	}

	// PokéAPI - Get Pokemon info
	fmt.Println("\n--- PokéAPI GET Example ---")
	resp, err = client.Get(POKEMON_API + "/pokemon/pikachu")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Pokemon Response: %s\n", string(body))
	}
}

// POST Examples
func postExamples(client *http.Client) {
	// JSONPlaceholder - Create post
	fmt.Println("\n--- JSONPlaceholder POST Example ---")
	post := Post{
		UserID: 1,
		Title:  "Test Post",
		Body:   "This is a test post",
	}
	jsonPost, _ := json.Marshal(post)
	resp, err := client.Post(JSONPLACEHOLDER_API+"/posts",
		"application/json",
		bytes.NewBuffer(jsonPost))
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Create Post Response: %s\n", string(body))
	}
}

// PUT Example
func putExample(client *http.Client) {
	fmt.Println("\n--- JSONPlaceholder PUT Example ---")
	post := Post{
		UserID: 1,
		Title:  "Updated Post",
		Body:   "This is an updated post",
	}
	jsonPost, _ := json.Marshal(post)

	req, _ := http.NewRequest(http.MethodPut,
		JSONPLACEHOLDER_API+"/posts/1",
		bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Update Post Response: %s\n", string(body))
	}
}

// DELETE Example
func deleteExample(client *http.Client) {
	fmt.Println("\n--- JSONPlaceholder DELETE Example ---")
	req, _ := http.NewRequest(http.MethodDelete,
		JSONPLACEHOLDER_API+"/posts/1",
		nil)

	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		fmt.Printf("Delete Status: %d\n", resp.StatusCode)
	}
}

// HTTPBin Examples for special cases
func httpbinExamples(client *http.Client) {
	// Test headers
	fmt.Println("\n--- HTTPBin Headers Example ---")
	req, _ := http.NewRequest("GET", HTTPBIN_API+"/headers", nil)
	req.Header.Add("Custom-Header", "test-value")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Headers Response: %s\n", string(body))
	}

	// Test cookies
	fmt.Println("\n--- HTTPBin Cookies Example ---")
	req, _ = http.NewRequest("GET", HTTPBIN_API+"/cookies", nil)
	req.AddCookie(&http.Cookie{Name: "test-cookie", Value: "cookie-value"})
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Cookies Response: %s\n", string(body))
	}

	// Test delay
	fmt.Println("\n--- HTTPBin Delay Example ---")
	resp, err = client.Get(HTTPBIN_API + "/delay/1")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Delay Response: %s\n", string(body))
	}
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Run examples
	getExamples(client)
	// postExamples(client)
	// putExample(client)
	// deleteExample(client)
	// httpbinExamples(client)
}
