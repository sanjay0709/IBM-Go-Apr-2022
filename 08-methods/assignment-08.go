package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%v, Expiry = %s", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}
func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    50,
		Category: "Stationary",
	}

	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	//Using PerishableProduct
	grapes := NewPerishableProduct(200, "Grapes", 40, 10, "Fruits", "2 Days")
	fmt.Println(grapes.Format())

	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())

}
