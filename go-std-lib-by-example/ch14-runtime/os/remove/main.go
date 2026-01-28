package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("demo.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	*/

	err := os.Remove("demo.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted successfully.")
	}

	/*
		dirs := "dir1/dir2/ird3"
		err = os.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	*/

	// err = os.Remove("dir1")  // Error, Dir not empty
	// /*
	err = os.RemoveAll("dir1")
	if err != nil {
		fmt.Println("Error deleting directory:", err)
	} else {
		fmt.Println("Directory deleted successfully.")
	}
	// */
}
