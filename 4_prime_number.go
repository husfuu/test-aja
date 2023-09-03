package main

import (
	"fmt"
	"math"
)

func PrintPrimeNumber1_100() {
	for i := 0; i < 100; i++ {
		if IsPrimeNumber(i) {
			fmt.Printf("Bilangan prima: %d \n", i)
		} else if i%9 == 0 {
			kelipatan := i / 9
			fmt.Printf("Kelipatan 9 ke-%d: %d \n", kelipatan, i)
		}
	}
}

func IsPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}
	sqRoot := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqRoot; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
