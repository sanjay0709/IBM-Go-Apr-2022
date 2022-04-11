package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter the name:")
	fmt.Scanln(&name)
	fmt.Println("Hi ", name)

	var n1, n2 int
	fmt.Println("Enter 2 numbers")
	//fmt.Scanln(&n1, &n2)
	fmt.Scanf("%d-%d\n", &n1, &n2)
	fmt.Println(n1 + n2)
}
