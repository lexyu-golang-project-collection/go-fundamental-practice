package main

import (
	"fmt"
)

func main() {
Menu:
	for {
		fmt.Println("\n*******請選擇範例：*********")
		fmt.Println("1 表示 io.Reader 範例")
		fmt.Println("2 表示 io.ByteReader/ByteWriter 範例")
		fmt.Println("3 表示 ReadAt 範例")
		fmt.Println("4 表示 WriteAt 範例")
		fmt.Println("5 表示 ReaderFrom 範例")
		fmt.Println("6 表示 WriteTo 範例")
		fmt.Println("7 表示 Seeker 範例")
		fmt.Println("8 表示 LimitedReader 範例")
		fmt.Println("9 表示 Pipes 範例")
		fmt.Println("10 表示 Copy Std 範例")
		fmt.Println("11 表示 Copy N 範例")
		fmt.Println("12 表示 ReadAtLeast 範例")
		fmt.Println("13 表示 ReadFull 範例")
		fmt.Println("14 表示 MultiReader 範例")
		fmt.Println("15 表示 MultiWriter 範例")
		fmt.Println("16 表示 TeeReader 範例")
		fmt.Println("q 退出")
		fmt.Println("***********************************")

		var ch string
		fmt.Scanln(&ch)
		switch ch {
		case "1":
			ReaderExamples()
		case "2":
			ByteWriterExample()
		case "3":
			ReadAtDemo()
		case "4":
			WriteAtDemo()
		case "5":
			ReaderFromDemo()
		case "6":
			WriteToDemo()
		case "7":
			SeekerDemo()
		case "8":
			LimitedReaderDemo()
		case "9":
			PipesDemo()
		case "10":
			CopyStdDemo()
		case "11":
			CopyNDemo()
		case "12":
			ReadAtLeastDemo()
		case "13":
			ReadFullDemo()
		case "14":
			MultiReaderDemo()
		case "15":
			MultiWriterDemo()
		case "16":
			TeeReaderDemo()
		case "q":
			fmt.Println("exit !")
			break Menu
		default:
			fmt.Println("error !")
			continue
		}
	}
}
