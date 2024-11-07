package main

import (
	"fmt"
	"path"
)

func main() {
	// Join paths
	p1 := path.Join("..", "/csv/to_json/", "data.csv")
	fmt.Println("Joined Path:", p1)

	// Clean a path (e.g., remove redundant slashes or dots)
	p2 := path.Clean("/home/user/../documents/./file.txt")
	fmt.Println("Cleaned Path:", p2)

	// Get the base (last element) of the path
	p3 := path.Base("/home/user/documents/file.txt")
	fmt.Println("Base Path:", p3)

	// Get the directory of the path
	p4 := path.Dir("/home/user/documents/file.txt")
	fmt.Println("Directory Path:", p4)

	// Check if a path is absolute
	p51 := path.IsAbs("/home/user/documents/file.txt")
	fmt.Println("Is Absolute Path:", p51)
	p52 := path.IsAbs("home/user/documents/file.txt")
	fmt.Println("Is Absolute Path:", p52)

	// Get the extension of the file
	p6 := path.Ext("file.txt")
	fmt.Println("File Extension:", p6)
}
