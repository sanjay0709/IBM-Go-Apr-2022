package main

import (
	"fmt"
	/* calc "modularity-demo/calculator" // package alias */
	_ "Modularity-Demo/calculator" //to execute the "init" function
	"Modularity-Demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("modularity demo")
	/* fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("opCount = ", calc.OpCount()) */

	fmt.Println("Is 21 an even number ? : ", utils.IsEven(21))
	fmt.Println("Is 21 an odd number ? : ", utils.IsOdd(21))
	color.Red("This string will be printed in red color")
	color.Green("This string will be printed in green color")
}
