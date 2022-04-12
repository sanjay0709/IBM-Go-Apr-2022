package main

import "fmt"

func main() {
	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice > 5 || userChoice < 1 {
			fmt.Println("Invalid choice")
			continue
		}
		processUserChoice(userChoice)
	}
}

func processUserChoice(userChoice int) {
	result := 0
	n1, n2 := getInputs()
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	fmt.Println("Result = ", result)
}

func getUserChoice() (userChoice int) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter you choice:")
	fmt.Scanln(&userChoice)
	return
}

func getInputs() (n1, n2 int) {
	fmt.Println("Enter the nos :")
	fmt.Scanln(&n1, &n2)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
