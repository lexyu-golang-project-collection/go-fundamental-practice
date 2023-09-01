package mypackage

import (
	"fmt"
	"os"
	"strconv"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"
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

	g.PL("iArgs = ", iArgs)

	max := 0
	for _, val := range iArgs {
		// if val >= max {
		max += val
		// }
	}
	fmt.Println("Max Value :", max)
}
