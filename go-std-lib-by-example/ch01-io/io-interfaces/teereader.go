package main

import (
	"io"
	"os"
	"strings"
)

func TeeReaderDemo() {
	reader := io.TeeReader(strings.NewReader("Go Langauge Practice"), os.Stdout)
	reader.Read(make([]byte, 50))
}
