package main

import (
	"io"
	"os"
)

func MultiWriterDemo() {
	file, err := os.Create("../01/tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("Go Language Practice"))
}
