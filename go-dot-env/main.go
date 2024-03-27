package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("HOME = ", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Println("The env var SHELL is not set")
	} else {
		fmt.Println("SHELL = ", shell)
	}

	err := os.Setenv("TEST_NAME", "TESTNAME")
	if err != nil {
		fmt.Printf("Could not set the env var TEST_NAME")
	}
	fmt.Printf("TEST_NAME = %s\n", os.Getenv("TEST_NAME"))

	// GodotEnv

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	fmt.Printf("API_KEY = %s\n", os.Getenv("API_KEY"))
	fmt.Printf("DB_PASS = %s\n", os.Getenv("DB_PASS"))

	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		fmt.Printf("Error loading .env into map[string]string\n")
		os.Exit(1)
	}

	for k, v := range envMap {
		fmt.Printf("Key = %s, Value = %s\n", k, v)
	}
	fmt.Printf("API_KEY = %s\n", envMap["API_KEY"])
}
