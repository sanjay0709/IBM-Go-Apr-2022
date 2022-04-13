package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
	Create a PerishableProduct with "Expiry (string)" field (and all the other fields of Product)
	Use the Format and ApplyDiscount functions for PerishableProduct
	Create a factory function "NewPerishableProduct" to create a new perishable product
*/

type PerishableProduct struct {
	Product
	Expiry string
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

	fmt.Println(Format(pen))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))

	//Using PerishableProduct
	grapes := NewPerishableProduct(200, "Grapes", 40, 10, "Fruits", "2 Days")
	fmt.Println(Format(grapes.Product))

	fmt.Println("After applying 10% discount")
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println(Format(grapes.Product))

}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
}
