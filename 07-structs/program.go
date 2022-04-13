package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{}
	//var emp Employee = Employee{100, "Magesh", "Kuppan", 10000}
	//var emp Employee = Employee{Id: 100, FirstName: "Magesh", Salary: 10000}

	/* var emp Employee = Employee{
		Id:        100,
		FirstName: "Magesh",
		Salary:    10000,
	} */

	var emp *Employee = new(Employee) // pointer to a newly created Employee object
	//fmt.Println(emp)
	fmt.Printf("%#v\n", emp)
	fmt.Println(emp.Id)

	/*
		e1 := Employee{200, "Suresh", "Kannan", 20000}
		e2 := Employee{200, "Suresh", "Kannan", 20000}
	*/
	e1 := new(Employee)
	e2 := new(Employee)
	fmt.Println(e1 == e2)

	n1 := 100
	n2 := 100
	fmt.Println(n1 == n2)
}
