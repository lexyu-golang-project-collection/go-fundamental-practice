package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 12) {
		return errors.New("incorrect month value")
	}
	d.month = m
	return nil
}

func (d *Date) SetYear(y int) error {
	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = y
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

func OOP() {
	date := Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day : %d/%d/%d\n", date.Month(), date.Day(), date.Year())
}

func main() {
	OOP()
}
