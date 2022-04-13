package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    50,
		Category: "Stationary",
	}

	fmt.Println(Format(pen))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))

	/*
		fmt.Println(pen.Cost)
		var penPtr *Product = &pen
		//fmt.Println((*penPtr).Cost)
		fmt.Println(penPtr.Cost)
	*/

}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
}
