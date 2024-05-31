package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println(rand.Intn(100))
		x++
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(rand.Intn(1000))

}
