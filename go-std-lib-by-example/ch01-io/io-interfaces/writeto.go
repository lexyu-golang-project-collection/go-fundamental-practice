package main

import (
	"bytes"
	"os"
)

func WriteToDemo() {
	reader := bytes.NewReader([]byte("這是一段句子"))
	reader.WriteTo(os.Stdout)
}
