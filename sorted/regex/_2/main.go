package main

import (
	"fmt"
	"regexp"
)

const (
	STR  = "12345"
	STR2 = "A124"
	STR3 = "94C"
	STR4 = ""
	STR5 = "Z"
	STR6 = "123456"
	STR7 = "1234567"
)

func main() {
	strs := []string{STR, STR2, STR3, STR4, STR5, STR6, STR7}
	reg := regexp.MustCompile(`^[A-Za-z0-9]{1,5}`)
	reg2 := regexp.MustCompile(`^[A-Za-z0-9]{1,5}$`)

	for i, v := range strs {
		if reg.Match([]byte(v)) {
			fmt.Println("reg ", i+1, v, true)
		} else {
			fmt.Println("reg ", i+1, v, false)
		}
		if reg2.Match([]byte(v)) {
			fmt.Println("reg2 ", i+1, v, true)
		} else {
			fmt.Println("reg2 ", i+1, v, false)
		}
		fmt.Println("////////////////////////////")
	}
}
