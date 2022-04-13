package main

import "fmt"

func main() {
	//var x interface{} //before go 1.18
	var x any // from go 1.18
	//x = 100
	//x = "a sample string"
	//x = true
	//x = 10.05
	//x = struct{}{}
	x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	//y := x.(int) + 100

	if val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an integer, x + 200 = ", val+200)
	case string:
		fmt.Println("x is a string, len(s) = ", len(val))
	case float64:
		fmt.Println("x is a float, x + 200.50 = ", val+200.50)
	case bool:
		fmt.Println("x is a boolean, value = ", val)
	case struct{}:
		fmt.Println("x is an empty struct")
	case []int:
		fmt.Println("x is a slice of int : ", val)
	default:
		fmt.Println("Unknown type")
	}

}
