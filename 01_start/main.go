package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch year := year; month {
	case 2:
		if year % 4 == 0 {
			daysInMonth = 28
		} else {
			if year%100 == 100 {
				if year%400 == 0 {
					daysInMonth = 28
				} else {
					daysInMonth = 29
				}
			} else {
				daysInMonth = 29
			}
		}

	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
