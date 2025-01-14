package main

import (
	"fmt"
)

// 4. 泛型方式 (使用自定義約束)
type customConstraint interface {
	~int | ~string
}

type MyInt int
type MyString string

func getKeysCustom[K customConstraint, V any](m map[K]V) []K {
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

	mapMyIntFloat := map[MyInt]float64{
		MyInt(1): 1.1,
		MyInt(2): 2.2,
	}

	fmt.Println("\n4. 泛型方式 (自定義約束):")
	fmt.Printf("string keys: %v\n", getKeysCustom(mapStringInt))
	fmt.Printf("int keys: %v\n", getKeysCustom(mapIntString))
	fmt.Printf("MyInt keys: %v\n", getKeysCustom(mapMyIntFloat))

	// Compilation Error Example
	// mapFloatInt := map[float64]int{1.1: 1, 2.2: 2}
	// fmt.Printf("float keys: %v\n", getKeysCustom(mapFloatInt)) // 錯誤：float64 不符合 customConstraint
}
