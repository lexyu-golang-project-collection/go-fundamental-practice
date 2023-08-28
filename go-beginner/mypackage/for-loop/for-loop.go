package mypackage

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"
)

func for_loop() {
	// for initialization; condition; postStatemnt{Body}
	for x := 1; x <= 5; x++ {
		g.PL(x)
	}

	fx := 0
	for fx < 5 {
		g.PL(fx)
		fx++
	}

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for {
		fmt.Print("Guess a number between 0 and 50 :")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			g.PL("Pick a Lower Value")
		} else if iGuess < randNum {
			g.PL("Pick a Higher Value")
		} else {
			g.PL("You Guessed it")
			g.PL("Random Number is :", randNum)
			break
		}
	}
}
