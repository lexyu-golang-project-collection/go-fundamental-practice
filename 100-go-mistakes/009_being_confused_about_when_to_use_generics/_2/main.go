package main

import (
	"fmt"
)

// 2. interface{} 方式
func getKeysInterface(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	}
}

func main() {
	mapStringInt := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	mapIntString := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println("\n2. interface{} 方式:")
	keys2, _ := getKeysInterface(mapStringInt)
	fmt.Printf("string keys: %v\n", keys2)
	keys2, _ = getKeysInterface(mapIntString)
	fmt.Printf("int keys: %v\n", keys2)
	// Err Example
	keys2, err := getKeysInterface([]string{"not", "a", "map"})
	fmt.Printf("error handling: %v\n", err)
}
