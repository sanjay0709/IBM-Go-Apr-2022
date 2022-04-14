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
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
		case ch <- "from main":
			time.Sleep(500 * time.Millisecond)
		default:
			fmt.Println("No channel operation were successful")
		}
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
