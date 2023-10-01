package main

import "fmt"

func calculate(a int, b int) (results []float64) {
	results = append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
