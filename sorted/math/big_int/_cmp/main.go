package main

import (
	"fmt"
	"math/big"
)

func calculateBigIntMemoryUsage(x *big.Int) int {
	bits := x.Bits()
	bytesUsed := len(bits) * (64 + 7) / 8 // 64 bits per big.Word, rounded up to the nearest byte

	return bytesUsed
}
func main() {

	big_a := big.NewInt(0)
	big_b := big.NewInt(1)

	// 將 limit 初始化為 10^99，即 100 digits 的最小整數。
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for big_a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		big_a.Add(big_a, big_b)
		// Swap a and b so that b is the next number in the sequence.
		big_a, big_b = big_b, big_a
	}
	fmt.Println(big_a) // 100-digit Fibonacci number

	fmt.Println("Length : ", len(big_a.Text(10)))
	memoryUsage := calculateBigIntMemoryUsage(big_a)
	fmt.Printf("Estimated memory usage of big_a: %v bytes\n", memoryUsage)
}
