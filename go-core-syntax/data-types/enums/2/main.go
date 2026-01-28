package main

import (
	"errors"
	"fmt"
	"log"
)

type Bank int

const (
	BKTW Bank = iota
	ESUN
	CTBC
)

type bank struct {
	code string
	name string
}

var banks = map[Bank]bank{
	BKTW: {code: "004", name: "臺灣銀行"},
	ESUN: {code: "808", name: "玉山銀行"},
	CTBC: {code: "822", name: "中國信託"},
}

func (b Bank) Code() string {
	return banks[b].code
}

func (b Bank) Name() string {
	return banks[b].name
}

func All() []Bank {
	return []Bank{BKTW, ESUN, CTBC}
}

func GetBank(code string) (Bank, error) {
	for _, b := range All() {
		if b.Code() == code {
			return b, nil
		}
	}
	return -1, errors.New("No match")
}

func main() {
	b, err := GetBank("808")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("(%s)%s\n", b.Code(), b.Name())
}
