package main

import "fmt"

func main() {
	/*
		nums := [5]int{1, 2, 3, 4, 5}

		for i, num := range nums {
			fmt.Println(i, num)
		}
	*/
	/*
		var wg sync.WaitGroup
		wg.Add(2)
		ch := make(chan int)
		go sell(ch, &wg)
		go buy(ch, &wg)
		wg.Wait()
	*/

	ch := make(chan int, 5)
	ch <- 100
	ch <- 200
	// close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}

/*
func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	fmt.Println("Sent all data to the channel")
	close(ch)
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	for val := range ch {
		fmt.Println("Received data : ", val)
	}
	// fmt.Println("Received data : ", <-ch)
	wg.Done()
}
*/
