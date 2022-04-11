package main

import "fmt"

func main() {

	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := can be used ONLY in a function (NOT at package level)
	fmt.Printf("type of msg = %T\n", msg)
	fmt.Println(msg)

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of x and y is "
		result = x + y
		fmt.Println(str, result)
	*/
	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of x and y is "
		var result int = x + y
		fmt.Println(str, result)
	*/
	/*
		var x, y int
		var str string
		x = 100
		y = 200
		str = "Sum of x and y is "
		var result int = x + y
		fmt.Println(str, result)
	*/
	/*
		var x, y int = 100, 200
		var str string = "Sum of x and y is "
		var result int = x + y
		fmt.Println(str, result)
	*/
	/*
		var x, y = 100, 200
		var str = "Sum of x and y is "
		var result = x + y
		fmt.Println(str, result)
	*/
	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of x and y is "
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of x and y is "
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var x, y, str = 100, 200, "Sum of x and y is "
		var result = x + y
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of x and y is "
	result := x + y
	fmt.Println(str, result)
}
