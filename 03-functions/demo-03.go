//anonymous functions

package main

import "fmt"

func main() {

	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}(100, 200)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing %d and %d, quotient = %d and remainder = %d\n", 100, 200, q, r)
}
