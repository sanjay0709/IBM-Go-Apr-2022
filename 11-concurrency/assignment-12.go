package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes()
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes() <-chan int {
	ch := make(chan int)
	//timeOutCh := timeout()
	timeOutCh := time.After(10 * time.Second)

	go func() {
		no := 3
	LOOP:
		for {
			if !isPrime(no) {
				no++
				continue LOOP
			}
			select {
			case <-timeOutCh:
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
