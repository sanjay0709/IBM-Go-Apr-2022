package main

import "fmt"

func main() {
	//if
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd nuimber\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd nuimber\n", no)
	}

	//for
	/* for - ver 1.0 */
	fmt.Println("for-loop (regular)")
	for x, no := 1, 100; no <= 10; no++ {
		fmt.Println(no, x)
	}

	/* for - ver 2.0 */
	fmt.Println("for-loop (while)")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	/* for - ver 3.0 */
	fmt.Println("for-loop (infinite)")
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Printf("x = %d\n", x)

	//switch
}
