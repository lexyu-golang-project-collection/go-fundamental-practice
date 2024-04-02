package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Gopher"

	encoded := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(encoded)
}
