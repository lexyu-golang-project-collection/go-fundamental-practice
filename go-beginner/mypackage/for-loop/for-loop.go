package mypackage

import (
	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"
)

func ForLoop() {
	// for initialization; condition; postStatemnt{Body}
	for x := 1; x <= 5; x++ {
		g.PL(x)
	}
	g.PL("------------------------------------------------------------")

	fx := 0
	for fx < 5 {
		g.PL(fx)
		fx++
	}

	g.PL("------------------------------------------------------------")

	/*
		seedSecs := time.Now().Unix()
		rand.Seed(seedSecs)
		randNum := rand.Intn(50) + 1
		for {
			g.PL("Guess a number between 0 and 50 :")
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
	*/

}
