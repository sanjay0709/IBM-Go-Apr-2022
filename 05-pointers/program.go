package main

import "fmt"

func main() {
	var no int = 100

	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//accessing underlying value from a pointer (address) => dereferencing
	var no2 int = *noPtr
	fmt.Println(no2)

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	x, y := 10, 20
	fmt.Printf("Before swapping x = %d and y = %d\n", x, y)
	swap( /*  */ )
	fmt.Printf("After swapping x = %d and y = %d\n", x, y)
}

func increment(x *int) {
	*x = *x + 1
}

func swap( /*  */ ) {

}
