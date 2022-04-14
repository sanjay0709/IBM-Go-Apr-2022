package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := f1()
	ch2 := f2()
	ch := make(chan string)
	go fn(ch)

	for {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
		ch <- "from main"
		time.Sleep(500 * time.Millisecond)
	}
}

func fn(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func f1() <-chan string {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ch <- "from f1"
		}
	}()
	return ch
}

func f2() <-chan string {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch <- "from f2"
		}
	}()
	return ch
}
