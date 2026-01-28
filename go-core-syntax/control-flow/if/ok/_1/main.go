package main

import "fmt"

func IsEven(n int) (bool, error) {
	if n&1 == 1 {
		return false, fmt.Errorf("it's an odd number")
	}
	return true, nil
}

func main() {

	if ok, err := IsEven(3); ok {
		fmt.Println("it's an even number")
	} else {
		fmt.Println(err)
	}

	if ok, err := IsEven(4); ok {
		fmt.Println("it's an even number")
	} else {
		fmt.Println(err)
	}

}
