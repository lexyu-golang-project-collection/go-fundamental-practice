package main

import (
	"crypto/md5"
	"fmt"
)

func MD5_1(str string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x\n", sum)
}

func main() {
	str := "Gopher"
	res := MD5_1(str)
	fmt.Println(res)
}
