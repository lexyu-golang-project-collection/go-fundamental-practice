package main

import (
	"fmt"
)

// 3. 泛型方式 (使用 comparable)
func getKeysGeneric[K comparable, V any](m map[K]V) []K {
	var keys []K
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

	mapIntString := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println("\n3. 泛型方式 (comparable):")
	fmt.Printf("string keys: %v\n", getKeysGeneric(mapStringInt))
	fmt.Printf("int keys: %v\n", getKeysGeneric(mapIntString))
	// accept any comparable 型別
	mapFloatBool := map[float64]bool{1.1: true, 2.2: false}
	fmt.Printf("float keys: %v\n", getKeysGeneric(mapFloatBool))
}
