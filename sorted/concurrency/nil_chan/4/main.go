package main

func main() {
	var ch chan int
	select {
	case <-ch: // This line will cause a deadlock
	}
}
