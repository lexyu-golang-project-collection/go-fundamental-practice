package main

import "fmt"

func main() {
	// UTF-8 編碼的位元組
	bytes := []byte{228, 184, 150} // '世' 的 UTF-8 編碼

	fmt.Println("十進制    十六進制")
	fmt.Println("-------------------")
	for _, b := range bytes {
		fmt.Printf("%d        %X\n", b, b)
	}

	// 更詳細的展示
	fmt.Println("\n完整展示：")
	fmt.Printf("位元組值（十進制）: [%d %d %d]\n", bytes[0], bytes[1], bytes[2])
	fmt.Printf("位元組值（十六進制）: [%X %X %X]\n", bytes[0], bytes[1], bytes[2])
}
