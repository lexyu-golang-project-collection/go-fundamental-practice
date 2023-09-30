package main

import "fmt"

func main() {
	chan1 := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			chan1 <- true
		}()
	}

	fmt.Println("-------------------------------------------")

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-chan1
	}

	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			chan1 <- true
		}(v)
	}

	fmt.Println("-------------------------------------------")

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-chan1
	}

	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}

	fmt.Println("-------------------------------------------")

	for i := 1; i <= 3; i++ {
		j := i
		prints = append(prints, func() { fmt.Println(j) })
	}
	for _, print := range prints {
		print()
	}

}
