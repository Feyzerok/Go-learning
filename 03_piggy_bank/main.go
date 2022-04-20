package main

import (
	"fmt"
	"math/rand"
)

const goal = 20.00

func main() {
	count := 0
	for i := 0.0; i < goal; {
		switch rand.Intn(3) {
		case 0:
			i = i + 0.05
		case 1:
			i = i + 0.10
		case 2:
			i = i + 0.25
		}
		fmt.Printf("%4.2f\n", i)
		count = count + 1
	}
	fmt.Printf("Моент в Копилке - %v", count)
}
