package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDirDemo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the directory path: ")
	scanner.Scan()
	dir := scanner.Text()

	if dir == "" {
		fmt.Println("No directory path provided.")
		return
	}
	listAll(dir, 0)
}

func listAll(path string, curHier int) {
	// fileInfos, err := ioutil.ReadDir(path)
	fileInfos, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		for tmpHeir := curHier; tmpHeir > 0; tmpHeir-- {
			fmt.Printf("|\t")
		}
		if info.IsDir() {
			for tmpHeir := curHier; tmpHeir > 0; tmpHeir-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {

			fmt.Println(info.Name())
		}
	}
}
