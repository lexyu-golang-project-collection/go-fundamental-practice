package mypackage

import (
	"fmt"
	"os"
	"strconv"
)

func CLI() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	iArgs := []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	pl("iArgs = ", iArgs)

	max := 0
	for _, val := range iArgs {
		// if val >= max {
		max += val
		// }
	}
	fmt.Println("Max Value :", max)
}
