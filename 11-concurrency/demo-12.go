package main

import "fmt"

//non-buffered channel

/*
func main() {
	ch := make(chan int)
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
*/

/*
func main() {
	ch := make(chan int)
	data := <-ch
	ch <- 100
	fmt.Println(data)
}
*/

/*
func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
*/

//buffered channel
func main() {
	ch := make(chan int, 1)
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
