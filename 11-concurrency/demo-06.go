package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("main started")
	wg.Add(1)  //initializing the counter to 1
	go f1(&wg) //scheduling this function for execution
	f2()
	wg.Wait() //blocked until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrements the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
