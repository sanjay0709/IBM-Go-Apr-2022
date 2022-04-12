package main

import "fmt"

func main() {
	//var productRanks map[string]int

	//var productRanks map[string]int = make(map[string]int)
	//productRanks["Pen"] = 4

	//make(map[string]int) === map[string]int{}

	//var productRanks map[string]int = map[string]int{"Pen": 4, "Pencil": 1, "Marker": 2}
	var productRanks map[string]int = map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 2,
	}
	fmt.Println(productRanks)

	fmt.Println("Adding a new item")
	productRanks["Mouse"] = 6
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%s] = %d\n", key, val)
	}

	fmt.Println("Checking the existence of a key")
	var keyToCheck = "Tablet"
	if value, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Key - %q exists with value - %d\n", keyToCheck, value)
	} else {
		fmt.Printf("Key - %q doesnot exist\n", keyToCheck)
	}

	fmt.Println("Removing an item")
	delete(productRanks, "Pen")
	fmt.Println("After removing Pen, productRanks = ", productRanks)
}
