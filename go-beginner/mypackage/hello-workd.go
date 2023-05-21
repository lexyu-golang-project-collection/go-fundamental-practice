package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lexyu-golang-project-collection/logger"
)

var pl = fmt.Println
var pf = fmt.Printf

func Hello() {
	fmt.Println("Hello Go!")

	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

	logger.Log("go get from public repo")
}
