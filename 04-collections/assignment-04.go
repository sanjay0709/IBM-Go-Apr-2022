package main

import (
	"fmt"
)

func main() {
	var primeNos = generatePrimes(3, 100)
	fmt.Println(primeNos)

	//refactor the following into the generatePrimes functions
	for no := 3; no <= 100; no++ {
		isPrime := true
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is prime\n", no)
		}
	}
}
