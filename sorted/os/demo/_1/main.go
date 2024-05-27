package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func createDataTXT() (*os.File, error) {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return f, err
}

func openDataTXT() (*os.File, error) {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f, err
}

func readDataTxt(f *os.File) {
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		fmt.Println("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
func appendNewDataInTxt() {
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File Doesn't Exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
func main() {

	f, _ := createDataTXT()

	f, _ = openDataTXT()

	defer f.Close()

	readDataTxt(f)

	// appendNewDataInTxt()
}
