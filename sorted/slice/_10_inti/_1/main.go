package main

import (
	"fmt"
	"time"
)

func main() {
	// 1st Choice
	var a []int
	timeStart := time.Now()

	for i := 1; i <= 1_000_000; i++ {
		a = append(a, i)
	}

	fmt.Printf("%+v, ", a)

	fmt.Println(time.Since(timeStart))
}
