package main

import "fmt"

func main() {
	arr3 := [2][2]int{
		{5, 6},
		{7, 8},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("arr3[%d][%d] = %d\n", i, j, arr3[i][j])
		}
	}
}
