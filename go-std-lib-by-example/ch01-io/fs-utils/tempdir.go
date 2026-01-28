package main

import (
	"fmt"
	"os"
)

func TempDirDemo() {
	// work, err = ioutil.TempDir("", "go-build")
	tmpdir, err := os.MkdirTemp("../1-2", "go-build-*")
	if err != nil {
		fmt.Println("Error creating temp directory:", err)
		return
	}
	// defer os.RemoveAll(tmpdir)

	fmt.Println("Temporary directory created:", tmpdir)

	tempFilePath := tmpdir + "/tempfile.txt"
	err = os.WriteFile(tempFilePath, []byte("TempDir: Hello, Go!"), 0644)
	if err != nil {
		fmt.Println("Error creating file in temp directory:", err)
		return
	}
	fmt.Println("File created in temp directory:", tempFilePath)
}
