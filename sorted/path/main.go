package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	relativePath := filepath.Join("../csv/to_json/", "data.csv")
	fmt.Println(relativePath)

	abPath, _ := filepath.Abs(relativePath)
	fmt.Println(abPath)

}
