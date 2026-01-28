package main

import (
	"fmt"
	"time"
)

func main() {
	// 3rd Choice
	z := make([]int, 1_000_000)
	timeStart := time.Now()

	for i := 0; i < 1_000_000; i++ {
		z = append(z, i)
	}

	fmt.Printf("%+v, ", z)

	fmt.Println(time.Since(timeStart))
}
