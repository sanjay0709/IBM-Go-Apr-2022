package main

import "fmt"

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

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
