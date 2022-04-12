/* Handling errors */
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("cannot divide by zero")

func main() {
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Cannot divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong")
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		err := DivideByZeroError
		return 0, err
	}
	result := x / y
	return result, nil
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
