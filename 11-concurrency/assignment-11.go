package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes()
	for i := 0; i < 10; i++ {
		primeNo := <-ch
		fmt.Println(primeNo)
	}
}

func genPrimes() <-chan int {
	ch := make(chan int)
	go func() {
		counter := 10
		no := 3
		for {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
				counter--
			}
			if counter <= 0 {
				break
			}
			no++
		}
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
