package main

import "fmt"

func init() {
	fmt.Println("REDIS")
}

func Store(key, value string) (k, v string) {
	return key, value
}
