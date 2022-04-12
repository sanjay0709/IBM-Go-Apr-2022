/* higher order functions - return functions as return values */

package main

import "fmt"

func main() {
	fn := getFn()
	fn()

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	logAdd := getLogOperation(add)
	logSubtract := getLogOperation(subtract)

	logAdd(100, 200)
	logSubtract(100, 200)
}

/* func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("invocation started")
	oper(x, y)
	fmt.Println("invocation completed")
} */

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		oper(x, y)
		fmt.Println("invocation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
