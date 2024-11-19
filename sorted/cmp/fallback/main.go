package main

import (
	"cmp"
	"fmt"
	"os"
)

func main() {
	port := cmp.Or(
		getPortFromEnv(),
		getPortFromFlag(),
		"8080",
	)

	fmt.Printf("Starting server on port: %s\n", port)
}

func getPortFromFlag() string {
	return "9991"
}

func getPortFromEnv() string {
	return os.Getenv("PORT")
}
