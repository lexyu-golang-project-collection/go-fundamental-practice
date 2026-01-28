package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("float to int ============================================================================>")
	floatVal := 1.5
	floatToIntVal := int(floatVal)
	fmt.Println(floatToIntVal)

	fmt.Println("string to int ============================================================================>")

	strVal := "5000000"
	strToIntVal, err := strconv.Atoi(strVal)
	fmt.Println(strToIntVal, err, reflect.TypeOf(strToIntVal))

	fmt.Println("int to string ============================================================================>")
	intVal := 5000000
	intToStrVal := strconv.Itoa(intVal)
	fmt.Println(intToStrVal)

	fmt.Println("string to float ============================================================================>")
	strVal2 := "3.14"
	if strToFloatVal, err := strconv.ParseFloat(strVal2, 64); err == nil {
		fmt.Println(strToFloatVal)
	} else {
		fmt.Println("Wrong")
	}

	fmt.Println("float to string ============================================================================>")
	floatToStrVal := fmt.Sprintf("%f", 3.14)
	fmt.Println(reflect.TypeOf(floatToStrVal), floatToStrVal)
}
