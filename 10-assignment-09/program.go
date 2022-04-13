package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, prod := range products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, prod := range products {
		if predicate(prod) {
			result = append(result, prod)
		}
	}
	return result
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, prod := range products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, prod := range products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

//using the sort.Sort() function

/* //sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//to sort the products by name
type ByName struct {
	Products
}

//overriding the 'Products.Less' method
func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//to sort the products by cost
type ByCost struct {
	Products
}

//overriding the 'Products.Less' method
func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//to sort the products by units
type ByUnits struct {
	Products
}

//overriding the 'Products.Less' method
func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//to sort the products by category
type ByCategory struct {
	Products
}

//overriding the 'Products.Less' method
func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	default:
		sort.Sort(products)
	}
}
*/

//using sort.Slice() function
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	default:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	stove := Product{102, "Stove", 5000, 5, "Utencil"}

	//IndexOf
	fmt.Println("Index of Stove = ", products.IndexOf(stove))

	//Includes
	fmt.Println("Includes (stove) = ", products.Includes(stove))

	//Filter

	fmt.Println()
	fmt.Println("Filter")

	fmt.Println("Costly Products")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)

	fmt.Println("Costly stationary products")
	costlyStationaryProductPredicate := func(product Product) bool {
		return costlyProductPredicate(product) && stationaryProductPredicate(product)
	}
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	fmt.Println(costlyStationaryProducts)

	//All
	fmt.Println("All")
	fmt.Println("Are all the products stationary products? : ", products.All(stationaryProductPredicate))

	//Any
	fmt.Println("Any")
	fmt.Println("Are there any stationary products? : ", products.Any(stationaryProductPredicate))

	fmt.Println()
	fmt.Println("Sorting")
	fmt.Println("Default Sort (by Id)")
	//sort.Sort(products) // will work because Products implements the sort.Interface interface
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort by Name")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	//sort.Sort(ByCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort by Units")
	//sort.Sort(ByUnits{products})
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort by Category")
	//sort.Sort(ByCategory{products})
	products.Sort("Category")
	fmt.Println(products)

}
