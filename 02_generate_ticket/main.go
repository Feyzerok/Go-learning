package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 60 * 24 * 60

func main() {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed := rand.Intn(15) + 16
		duration := distance / speed / secondsPerDay
		price := 20 + speed

		if rand.Intn(2) == 0 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}
		fmt.Printf("%-16v %4v %-10v $%4v\n", company, duration, trip, price)
	}
}
