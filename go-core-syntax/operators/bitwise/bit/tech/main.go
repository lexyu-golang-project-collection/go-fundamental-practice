package main

import "fmt"

func main() {
	// 測試數字
	n := 42 // 二進制：101010
	fmt.Printf("測試數字 n = %d (二進制：%06b)\n", n, n)

	// 1. 判斷奇偶
	isEven := func(x int) bool {
		return (x & 1) == 0
	}
	fmt.Printf("\n1. 判斷奇偶：\n")
	fmt.Printf("   %d 是否為偶數：%v\n", n, isEven(n))
	fmt.Printf("   %d 是否為偶數：%v\n", 7, isEven(7))

	// 2. 取得第 i 位
	getBit := func(x, i int) int {
		return (x >> i) & 1
	}
	fmt.Printf("\n2. 取得各個位元：\n")
	for i := 0; i < 6; i++ {
		fmt.Printf("   第 %d 位：%d\n", i, getBit(n, i))
	}

	// 3. 設定第 i 位為 1
	setBit := func(x, i int) int {
		return x | (1 << i)
	}
	pos := 2
	fmt.Printf("\n3. 設定第 %d 位為 1：\n", pos)
	fmt.Printf("   原始：%06b\n", n)
	fmt.Printf("   結果：%06b\n", setBit(n, pos))

	// 4. 設定第 i 位為 0
	clearBit := func(x, i int) int {
		return x & ^(1 << i)
	}
	fmt.Printf("\n4. 設定第 %d 位為 0：\n", pos)
	fmt.Printf("   原始：%06b\n", n)
	fmt.Printf("   結果：%06b\n", clearBit(n, pos))

	// 5. 切換第 i 位
	toggleBit := func(x, i int) int {
		return x ^ (1 << i)
	}
	fmt.Printf("\n5. 切換第 %d 位：\n", pos)
	fmt.Printf("   原始：%06b\n", n)
	fmt.Printf("   結果：%06b\n", toggleBit(n, pos))

	// 6. 清除最右邊的 1
	clearRightmostOne := func(x int) int {
		return x & (x - 1)
	}
	fmt.Printf("\n6. 清除最右邊的 1：\n")
	fmt.Printf("   原始：%06b\n", n)
	fmt.Printf("   結果：%06b\n", clearRightmostOne(n))

	// 7. 取得最右邊的 1
	getRightmostOne := func(x int) int {
		return x & (-x)
	}
	fmt.Printf("\n7. 取得最右邊的 1：\n")
	fmt.Printf("   原始：%06b\n", n)
	fmt.Printf("   結果：%06b\n", getRightmostOne(n))

	// 額外演示：計算二進制中 1 的個數
	countOnes := func(x int) int {
		count := 0
		for x != 0 {
			x = x & (x - 1) // 清除最右邊的 1
			count++
		}
		return count
	}
	fmt.Printf("\n額外演示 - 計算二進制中 1 的個數：\n")
	fmt.Printf("   數字 %d (%06b) 中有 %d 個 1\n", n, n, countOnes(n))
}
