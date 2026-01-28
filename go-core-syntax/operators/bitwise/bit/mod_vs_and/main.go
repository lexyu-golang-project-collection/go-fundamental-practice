package main

import "fmt"

func main() {
	// 使用 capacity = 16 (2的4次方) 作為例子
	capacity := 16
	mask := capacity - 1 // 15 (二進制：1111)

	// 測試數值
	testValues := []int{35, 88, 100, 1234, 9527}

	fmt.Printf("capacity = %d (二進制: %04b)\n", capacity, capacity)
	fmt.Printf("mask = capacity-1 = %d (二進制: %04b)\n\n", mask, mask)

	fmt.Println("比較 mod (%) 和 AND (&) 運算:")
	fmt.Println("Value\t\tBinary\t\tMod\tAND\t是否相等")
	fmt.Println("-----\t\t------\t\t---\t---\t--------")

	for _, value := range testValues {
		modResult := value % capacity
		andResult := value & mask
		equal := modResult == andResult

		fmt.Printf("%d\t\t%08b\t%d\t%d\t%v\n",
			value,     // 原始值
			value,     // 二進制表示
			modResult, // mod 結果
			andResult, // AND 結果
			equal,     // 兩者是否相等
		)
	}

	// 特別展示 88 的 AND 運算過程
	value := 88
	fmt.Printf("\n特別展示 %d 的 AND 運算過程：\n", value)
	fmt.Printf("value     = %d (%08b)\n", value, value)
	fmt.Printf("mask      = %d (%08b)\n", mask, mask)
	fmt.Printf("result    = %d (%08b)\n", value&mask, value&mask)
}
