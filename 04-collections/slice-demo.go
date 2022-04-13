package main

import "fmt"

func main() {
	//using the make function
	var nos []int = make([]int, 5)
	/*
		nos[0] = 3
		nos[1] = 1
		nos[2] = 4
		nos[3] = 2
		nos[4] = 5
		fmt.Println(nos)
	*/
	nos = append(nos, 10)
	nos = append(nos, 20, 30)
	fmt.Println(nos)
	//var products []string
	/*
		var products []string = []string{"Pen", "Pencil", "Marker"}
		products[0] = "Ballpoint-Pen"
		fmt.Println(products)
	*/
	//var products []string
	/*
		var products []string = []string{"Pen", "Pencil", "Marker"}
		products = append(products, "Scribble-Pad")
		fmt.Println(products)
	*/

	//var products []string
	var products = make([]string, 0, 5)
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Pen")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Pencil")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Marker")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Mouse")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Mouse-Pad")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)
	products = append(products, "Charger")
	fmt.Printf("len = %d, cap = %d, products = %v\n", len(products), cap(products), products)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo :] => from lo to len-1
		[:hi] => from 0 to hi - 1
		[:] => copy of the slice
	*/

	fmt.Println("Slicing")
	fmt.Println("products[2:4] => ", products[2:4])
	fmt.Println("products[:4] => ", products[:4])
	fmt.Println("products[2:] => ", products[2:])

	newProducts := products[2:4]
	newProducts[0] = "Sketch-Pen"
	fmt.Println("newProducts => ", newProducts)
	fmt.Println("products => ", products)

	x := []string{}
	x = append(x, products[:3]...)
	x = append(x, products[3:]...)
	fmt.Println(x)
}
