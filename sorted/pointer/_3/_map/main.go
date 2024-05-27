package main

import (
	"fmt"
)

func modify(m map[string]int) {
	m["K"] = 75
}

func main() {

	// Maps
	codes := make(map[string]int)
	codes["A"] = 65
	codes["F"] = 70
	fmt.Println(codes)
	modify(codes)
	fmt.Println(codes)

}
