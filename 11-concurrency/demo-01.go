package main

import "fmt"

func main() {
	go f1() //scheduling this function for execution
	f2()
	// --> blocked
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
