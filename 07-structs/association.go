package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
	Org       *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	org := &Organization{
		Name: "IBM",
		City: "Bengaluru",
	}
	emp1 := Employee{
		Id:        100,
		FirstName: "Magesh",
		LastName:  "Kuppan",
		Salary:    10000,
		Org:       org,
	}
	emp2 := Employee{
		Id:        101,
		FirstName: "Rajesh",
		LastName:  "Pandit",
		Salary:    20000,
		Org:       org,
	}

	emp1.Org.City = "Pune"

	fmt.Println(emp1.Org.City)
	fmt.Println(emp2.Org.City)
}
