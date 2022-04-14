package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) //share memory by communicating
	fmt.Println("main started")
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan<- int) {
	result := x + y
	ch <- result
}
