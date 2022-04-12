package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("something went wrong. contact the administrator")
			fmt.Println(e)
			return
		}
		fmt.Println("application existing successfully")
	}()

	result, err := divideClient(100, 0)
	if err.Error() == "runtime error: integer divide by zero" {
		fmt.Println("Cannot divide by zero")
		return
	}
	if err != nil {
		fmt.Println("unknown error", err)
		return
	}
	fmt.Println(result)
	//?handle the panic here - NO
}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	result = divide(x, y)
	return
}

func divide(x, y int) int {
	result := x / y //=> panic occurs
	return result
}
