package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() //scheduling this function for execution
	f2()
	var input string
	fmt.Scanln(&input) // --> blocked (DO NOT USE THIS)
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
