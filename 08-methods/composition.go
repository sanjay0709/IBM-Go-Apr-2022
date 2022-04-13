package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

//methods for Employee
func (e Employee) Format() string {
	return fmt.Sprintf("Id = %d, FirstName = %q, LastName = %q, Salary = %v", e.Id, e.FirstName, e.LastName, e.Salary)
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

	emp := Employee{100, "Magesh", "Kuppan", 10000}
	fmt.Println(emp.Format())
	fte := NewFullTimeEmployee(101, "Rajesh", "Pandit", 10000, "Health Insurance")

	/*
		fmt.Println(fte.Employee.FirstName)
		fmt.Println(fte.FirstName)
	*/
	fmt.Println(fte.Employee.Format())
	fmt.Println(fte.Format())
}
