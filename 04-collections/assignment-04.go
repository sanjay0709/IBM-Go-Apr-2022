package main

import (
	"fmt"
)

func main() {
	var primeNos = generatePrimes(3, 100)
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	primeNos := make([]int, 0)
	for no := 3; no <= 100; no++ {
		if isPrime(no) {
			primeNos = append(primeNos, no)
		}
	}
	return primeNos
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
