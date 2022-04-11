package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		result = 0
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter you choice:")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice > 5 || userChoice < 1 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the nos :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result = ", result)
	}
}
