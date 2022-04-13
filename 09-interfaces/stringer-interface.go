package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

//methods for Employee
//implementation for the fmt.Stringer interface
func (e Employee) String() string {
	return fmt.Sprintf("Id = %d, FirstName = %q, LastName = %q, Salary = %v", e.Id, e.FirstName, e.LastName, e.Salary)
}

//function -> argument has to a 'pointer to employee'
func ChangeName(emp *Employee, fullName string) {
	names := strings.Split(fullName, " ")
	emp.FirstName = names[0]
	emp.LastName = names[1]
}

//the above as a method -> this can be invoked with a value OR a pointer
func (emp *Employee) ChangeName(fullName string) {
	names := strings.Split(fullName, " ")
	emp.FirstName = names[0]
	emp.LastName = names[1]
}

type FullTimeEmployee struct {
	Employee
	Benefits string
}

//method overriding
//implementation for the fmt.Stringer interface
func (fte FullTimeEmployee) String() string {
	//return fmt.Sprintf("Id = %d, FirstName = %q, LastName = %q, Salary = %v, Benefits = %q", fte.Id, fte.FirstName, fte.LastName, fte.Salary, fte.Benefits)
	return fmt.Sprintf("%v, Benefits = %q", fte.Employee, fte.Benefits)
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
	fmt.Println(emp)
	fte := NewFullTimeEmployee(101, "Rajesh", "Pandit", 10000, "Health Insurance")

	/*
		fmt.Println(fte.Employee.FirstName)
		fmt.Println(fte.FirstName)
	*/
	fmt.Println(fte.Employee)
	fmt.Println(fte)

	fmt.Println()
	fmt.Println("Change the name of the employee to [Ramesh Jayaraman]")
	//ChangeName(&emp, "Ramesh Jayaraman")
	//(&emp).ChangeName("Ramesh Jayaraman") // dont have to do this
	emp.ChangeName("Ramesh Jayraman")
	fmt.Println("After changing the name, emp => ")
	fmt.Println(emp)
}
