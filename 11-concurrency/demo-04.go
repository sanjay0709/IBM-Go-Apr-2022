package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(10) //initializing the counter to 1
	go f1()    //scheduling this function for execution
	f2()
	wg.Wait() //blocked until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrements the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
