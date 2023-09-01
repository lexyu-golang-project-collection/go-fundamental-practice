package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"os"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"
	"github.com/lexyu-golang-project-collection/logger"
)

func Hello() {
	fmt.Println("Hello Go!")

	g.PL("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		g.PL("Hello", name)
	} else {
		log.Fatal(err)
	}

	logger.Log("go get from public repo")
}

func Empty() {

}
