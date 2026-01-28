package main

import (
	"fmt"
)

func main() {
Menu:
	for {
		fmt.Println("\n*******請選擇範例：*********")
		fmt.Println("1 表示 ReadDir 範例")
		fmt.Println("2 表示 WriteAndReadFileDemo 範例")
		fmt.Println("3 表示 TempDir 範例")
		fmt.Println("4 表示 TempFile 範例")
		fmt.Println("5 表示 NopCloser 範例")
		fmt.Println("q 退出")
		fmt.Println("***********************************")

		var ch string
		fmt.Scanln(&ch)
		switch ch {
		case "1":
			ReadDirDemo()
		case "2":
			WriteAndReadFileDemo()
		case "3":
			TempDirDemo()
		case "4":
			TempFileDemo()
		case "5":
			NopCloserDemo()
		case "q":
			fmt.Println("exit !")
			break Menu
		default:
			fmt.Println("error !")
			continue
		}
	}
}
