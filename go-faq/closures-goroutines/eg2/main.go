package main

import (
	"fmt"
	"time"
)

func main() {

	// go 1.22 version fixed it
	names := []string{"Percy", "Gopher", "Santa"}

	for _, v := range names {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(2 * time.Second)
}
