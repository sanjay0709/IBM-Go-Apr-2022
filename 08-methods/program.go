package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

func (e Employee) WhoAmI() {
	fmt.Println("I am ", e.FirstName, e.LastName)
}

func main() {
	emp := Employee{
		Id:        100,
		FirstName: "Magesh",
		LastName:  "Kuppan",
		Salary:    100,
	}
	//WhoAmI(emp)
	emp.WhoAmI()
}
