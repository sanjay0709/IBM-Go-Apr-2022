/* defered functions */

package main

import "fmt"

func main() {
	/* defer func() {
		fmt.Println("	[defer] - main")
	}() */
	defer fmt.Println("	[defer] - main")
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		fmt.Println("	[defer-1] - f1")
	}()
	defer func() {
		fmt.Println("	[defer-2] - f1")
	}()
	defer func() {
		fmt.Println("	[defer-3] - f1")
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("	[defer] - f2")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
