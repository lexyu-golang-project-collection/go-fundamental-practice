package main

import (
	"fmt"
)

// Define custom type for permissions (bitmask)
type Permission uint

// Enum-like constants using bitmask
const (
	Read    Permission = 1 << iota // 1 << 0 = 1
	Write                          // 1 << 1 = 2
	Execute                        // 1 << 2 = 4
)

func main() {
	perm := Read | Write
	fmt.Println(perm) // Output: 3 (Read + Write)

	// Check if permission includes Write
	if perm&Write != 0 {
		fmt.Println("Has Write Permission")
	}
}
