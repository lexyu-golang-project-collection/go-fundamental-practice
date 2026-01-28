package main

import "fmt"

func main() {
	var m map[string]struct{}

	m = make(map[string]struct{})

	m["Red"] = struct{}{}
	m["Blue"] = struct{}{}

	_, ok := m["Red"]
	fmt.Println("Is Red in the map?", ok)

	_, ok = m["Green"]
	fmt.Println("Is Green in the map?", ok)
}
