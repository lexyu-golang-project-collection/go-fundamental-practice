package main

import (
	"fmt"
	"time"
)

func main() {
	// 2nd Choice
	b := make([]int, 0, 1_000_000)

	timeStart := time.Now()

	for i := 1; i <= 1_000_000; i++ {
		b = append(b, i)
	}

	fmt.Printf("%+v, ", b)

	fmt.Println(time.Since(timeStart))
}
