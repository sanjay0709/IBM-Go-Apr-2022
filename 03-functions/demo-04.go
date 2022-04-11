//higher order functions - functions can be assigned to variables
package main

import "fmt"

func main() {
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}"
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var greetUser func(string) string
	greetUser = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}
	fmt.Print(greetUser("Suresh"))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (q, r int) {
		q = x / y
		r = x % y
		return
	}
	fmt.Println(divide(100, 7))
}
