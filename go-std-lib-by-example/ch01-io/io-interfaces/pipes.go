package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func PipesDemo() {
	pipeReader, pipeWriter := io.Pipe()

	go PipeWriteDemo(pipeWriter)
	go PipeReadeDemo(pipeReader)

	time.Sleep(30 * time.Second)
}

func PipeWriteDemo(writer *io.PipeWriter) {
	data := []byte("GoStd練習中")
	for i := 0; i < 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("寫入字節 %d\n", n)
	}
	writer.CloseWithError(errors.New("PipeWriter已關閉"))
}

func PipeReadeDemo(reader *io.PipeReader) {
	buf := make([]byte, 128)

	for {
		fmt.Println("Receive Blocked 5 second")
		time.Sleep(5 * time.Second)
		fmt.Println("Receive Start")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字節: %d\n", n)
		fmt.Printf("buffer 内容: %s\n", buf)
	}
}
