package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("main started")
	for i := 1; i <= count; i++ {
		wg.Add(1) //initializing the counter to 1
		go f1(i)  //scheduling this function for execution
	}
	f2()
	wg.Wait() //blocked until the wg counter becomes 0
	fmt.Println("main completed")
	var input string
	fmt.Println("Hit ENTER to exit....")
	fmt.Scanln(&input)
}

func f1(i int) {
	fmt.Printf("f1 [%d] started\n", i)
	time.Sleep(5 * time.Second)
	fmt.Printf("f1 [%d] completed\n", i)
	wg.Done() // decrements the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
