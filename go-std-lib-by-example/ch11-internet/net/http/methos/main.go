package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// API endpoints
const (
	JSONPLACEHOLDER_API = "https://jsonplaceholder.typicode.com"
	HTTPBIN_API         = "https://httpbin.org"
	DOG_API             = "https://dog.ceo/api"
	DOG_IMAGE_API       = "https://dog.ceo/api/breeds/image/random"
)

// Global client
var Client = &http.Client{
	Timeout: 10 * time.Second,
}

// Basic response types
type Post struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type DogAPIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HTTP Methods Demo
func HttpMethodsDemo() {
	// GET
	fmt.Println("\n=== GET Example ===")
	resp, _ := Client.Get(JSONPLACEHOLDER_API + "/posts/1")
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET Response: %s\n", string(body))
	resp.Body.Close()

	// POST
	fmt.Println("\n=== POST Example ===")
	post := Post{UserID: 1, Title: "Test", Body: "Content"}
	jsonData, _ := json.Marshal(post)
	resp, _ = Client.Post(JSONPLACEHOLDER_API+"/posts",
		"application/json",
		bytes.NewBuffer(jsonData))
	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("POST Response: %s\n", string(body))
	resp.Body.Close()

	// PUT (使用 NewRequest)
	fmt.Println("\n=== PUT Example ===")
	req, _ := http.NewRequest(http.MethodPut,
		JSONPLACEHOLDER_API+"/posts/1",
		bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = Client.Do(req)
	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("PUT Response: %s\n", string(body))
	resp.Body.Close()

	// DELETE
	fmt.Println("\n=== DELETE Example ===")
	req, _ = http.NewRequest(http.MethodDelete,
		JSONPLACEHOLDER_API+"/posts/1", nil)
	resp, _ = Client.Do(req)
	fmt.Printf("DELETE Status: %d\n", resp.StatusCode)
	resp.Body.Close()
}

// Dog image functions
func SaveDogImage(directory string) error {
	// Get image URL
	resp, _ := Client.Get(DOG_IMAGE_API)
	var dogResp DogAPIResponse
	json.NewDecoder(resp.Body).Decode(&dogResp)
	resp.Body.Close()

	// Create directory
	os.MkdirAll(directory, 0755)

	// Download and save image
	resp, _ = Client.Get(dogResp.Message)
	defer resp.Body.Close()

	filename := filepath.Join(directory,
		fmt.Sprintf("dog_%d%s",
			time.Now().Unix(),
			filepath.Ext(dogResp.Message)))

	file, _ := os.Create(filename)
	defer file.Close()

	io.Copy(file, resp.Body)
	fmt.Printf("Image saved: %s\n", filename)
	return nil
}

func ShowDogAscii() error {
	// Get image URL
	resp, _ := Client.Get(DOG_IMAGE_API)
	var dogResp DogAPIResponse
	json.NewDecoder(resp.Body).Decode(&dogResp)
	resp.Body.Close()

	// Download and convert image
	resp, _ = Client.Get(dogResp.Message)
	defer resp.Body.Close()

	img, _, _ := image.Decode(resp.Body)
	printAsciiArt(img)
	return nil
}

func printAsciiArt(img image.Image) {
	chars := []string{"@", "%", "#", "*", "+", "=", "-", ":", ".", " "}
	bounds := img.Bounds()
	for y := 0; y < bounds.Max.Y; y += 8 {
		for x := 0; x < bounds.Max.X; x += 4 {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			gray := (r + g + b) / 3
			idx := int(gray/257) * len(chars) / 256
			if idx >= len(chars) {
				idx = len(chars) - 1
			}
			fmt.Print(chars[idx])
		}
		fmt.Println()
	}
}

func main() {
	// Demo HTTP methods
	// HttpMethodsDemo()

	// Demo dog image functions
	fmt.Println("\n=== Dog Image Examples ===")
	SaveDogImage("dogs")
	ShowDogAscii()
}
