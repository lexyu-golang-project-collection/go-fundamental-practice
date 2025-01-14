package main

import (
	"fmt"
)

// 1. 具體型別方式 (最原始)
func getKeysSpecific(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	mapStringInt := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	fmt.Println("1. 具體型別方式:")
	fmt.Printf("keys: %v\n", getKeysSpecific(mapStringInt))
}
