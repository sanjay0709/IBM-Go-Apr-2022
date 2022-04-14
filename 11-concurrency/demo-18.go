package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	ch := genPrimes(stop)

	go func() {
		var input string
		fmt.Scanln(&input)
		stop <- true
	}()

	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes(stop chan bool) <-chan int {
	ch := make(chan int)

	go func() {
		no := 3
	LOOP:
		for {
			if !isPrime(no) {
				no++
				continue LOOP
			}
			select {
			case <-stop:
				break LOOP
			case ch <- no:
				time.Sleep(500 * time.Millisecond)
				no++
			}
		}
		close(ch)
	}()
	return ch
}

func timeout() <-chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(20 * time.Second)
		timeOutCh <- time.Now()
	}()
	return timeOutCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
