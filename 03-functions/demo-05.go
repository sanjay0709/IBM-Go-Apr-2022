/* higher order functions - pass functions as arguments */

package main

import "fmt"

func main() {
	exec(fn)

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)

}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("invocation started")
	oper(x, y)
	fmt.Println("invocation completed")
}

/*
func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
}
*/

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	f()
}
