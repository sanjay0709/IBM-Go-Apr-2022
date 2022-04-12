package calculator

import "fmt"

var opCount int

//package init function
func init() {
	fmt.Println("calculator package initialized")
}

func OpCount() int {
	return opCount
}
