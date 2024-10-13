package main

import (
	"fmt"
	"time"
)

type RPC struct {
	result int
	done   chan struct{}
}

func (rpc *RPC) compute() {
	time.Sleep(time.Second)
	rpc.result = 100
	close(rpc.done)
}

func (RPC) version() int {
	return 1
}

func (rpc *RPC) version2() int {
	return 2
}

// go run -race main.go
func main() {
	// Demo1()
	Demo2()
}

func Demo1() {
	rpc := &RPC{done: make(chan struct{})}

	go rpc.compute()

	version := rpc.version()

	<-rpc.done
	result := rpc.result

	fmt.Printf("result: %d, version: %d\n", result, version)
}

func Demo2() {
	rpc := &RPC{done: make(chan struct{})}

	go rpc.compute()

	version2 := rpc.version2()

	<-rpc.done
	result := rpc.result

	fmt.Printf("result: %d, version2: %d\n", result, version2)
}
