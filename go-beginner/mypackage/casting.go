package mypackage

import (
	"fmt"
	"reflect"
	"strconv"
)

func Casting() {
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println(cV2)

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	fmt.Println(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println(cV8)
	} else {
		fmt.Println("Wrong")
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cV9)
}
