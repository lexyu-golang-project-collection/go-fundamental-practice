package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("../demo.txt", "../new-demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	/* Directory
	err = os.Rename("olddir", "newdir")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	*/
}
