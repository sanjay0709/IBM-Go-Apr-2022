package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

type FullTimeEmployee struct {
	Employee
	Benefits string
}

func NewFullTimeEmployee(id int, firstName string, lastName string, salary float32, benefits string) FullTimeEmployee {
	return FullTimeEmployee{
		Employee: Employee{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Salary:    salary,
		},
		Benefits: benefits,
	}
}

func main() {
	/*
		fte := FullTimeEmployee{
			Employee: Employee{Id: 100, FirstName: "Magesh", LastName: "Kuppan", Salary: 10000},
			Benefits: "Health Insurance",
		}
	*/
	fte := NewFullTimeEmployee(100, "Magesh", "Kuppan", 10000, "Health Insurance")
	fmt.Println(fte.Benefits)
	//fmt.Println(fte.Employee.FirstName)
	fmt.Println(fte.FirstName)
}
