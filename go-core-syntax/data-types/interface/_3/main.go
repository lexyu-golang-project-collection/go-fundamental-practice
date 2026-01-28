package main

import "fmt"

func main() {
	var s any = nil

	switch v := s.(type) {
	case string:
		fmt.Printf("string: %s", v)
	case int, uint:
		fmt.Printf("integer or unsigned integer: %v", v)
	case nil:
		fmt.Printf("nil value: %v", v)
	}
}
