package main

import (
	"fmt"
)

type MyStr string //alias for string

func (str MyStr) Length() int {
	return len(str)
}

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

	s := MyStr("This is a dummy string") // converting the "string" data to "MyStr" type
	fmt.Println(s.Length())

	//type convertion
	var f float32
	var i int8 = 100
	f = float32(i)
	fmt.Println(f)
}
