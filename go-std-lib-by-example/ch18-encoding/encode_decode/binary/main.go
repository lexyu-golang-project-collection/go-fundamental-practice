package main

import "fmt"

func main() {
	num := 5 // 二進位是 "101"

	fmt.Printf("%b\n", num)   // 輸出: "101"
	fmt.Printf("%8b\n", num)  // 輸出: "     101" (左邊補空格)
	fmt.Printf("%08b\n", num) // 輸出: "00000101" (左邊補0)

	// 用中文字 '世' 的 UTF-8 編碼示範
	buf := make([]byte, 3)
	buf[0] = 228 // 0xE4
	buf[1] = 184 // 0xB8
	buf[2] = 150 // 0x96

	for _, b := range buf {
		fmt.Printf("%08b ", b) // 每個位元組都顯示為 8 位二進位
	}
	// 輸出: 11100100 10111000 10010110
}
