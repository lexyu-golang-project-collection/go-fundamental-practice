package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lexyu-golang-project-collection/go-fundamental-practice/go-std-lib-by-example/util"
)

func ReaderMenu() {

	fmt.Println("\n*******從不同來源讀取資料*********")
	fmt.Println("*******請選擇資料來源，請輸入：*********")
	fmt.Println("1 表示 標準輸入")
	fmt.Println("2 表示 普通檔案")
	fmt.Println("3 表示 從字串")
	fmt.Println("4 表示 從網路")
	fmt.Println("b 返回上級選單")
	fmt.Println("q 退出")
	fmt.Println("***********************************")

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)

	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func ReaderExamples() {
	for {

		ReaderMenu()

		var str string
		fmt.Scanln(&str)

		var (
			data []byte
			err  error
		)

		switch strings.ToLower(str) {
		case "1":
			fmt.Println("輸入不超過12個字符")
			data, err = ReadFrom(os.Stdin, 12)
		case "2":
			file, err := os.Open(util.GetProjectRoot() + "/01/01.txt")
			if err != nil {
				fmt.Println("Open File Failed, ", err)
				continue
			}
			data, _ = ReadFrom(file, 12)
			file.Close()
		case "3":
			data, _ = ReadFrom(strings.NewReader("from string"), 12)
		case "4":
			fmt.Println("Pending...")
		case "b":
			fmt.Println("返回 Top Menu !")
			return
		case "q":
			fmt.Println("Quit !")
			os.Exit(0)
		}

		if err != nil {
			fmt.Println("err = ", err)
		} else {
			fmt.Println("讀取到的資料 = ", string(data))
		}
	}

}
