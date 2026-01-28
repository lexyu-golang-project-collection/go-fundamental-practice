package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	relativePath := filepath.Join("../../csv/to_json/", "data.csv")
	fmt.Println(relativePath)

	abPath, _ := filepath.Abs(relativePath)
	fmt.Println(abPath)

	// Clean a path (like path.Clean, but works platform-independently)
	p2 := filepath.Clean("/home/user/../documents/./file.txt")
	fmt.Println("Cleaned Path:", p2)

	// Get the base (last element) of the path
	p3 := filepath.Base("/home/user/documents/file.txt")
	fmt.Println("Base Path:", p3)

	// Get the directory of the path
	p4 := filepath.Dir("/home/user/documents/file.txt")
	fmt.Println("Directory Path:", p4)

	// Check if a path is absolute
	p5 := filepath.IsAbs("/home/user/documents/file.txt")
	fmt.Println("Is Absolute Path:", p5)

	// Get the file extension
	p6 := filepath.Ext("file.txt")
	fmt.Println("File Extension:", p6)

	// Get the relative path between two paths
	relPath, err := filepath.Rel("/home/user", "/home/user/documents/file.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Relative Path:", relPath)
	}

	// Example of walking through a directory using filepath.Walk
	relativePath = filepath.Join("../../csv")
	abPath, _ = filepath.Abs(relativePath)

	err = filepath.Walk(abPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Here info is of type os.FileInfo, not filepath.FileInfo
		fmt.Println("Visited file:", path)
		fmt.Println("Is directory:", info.IsDir())
		fmt.Println("File size:", info.Size())
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

}
