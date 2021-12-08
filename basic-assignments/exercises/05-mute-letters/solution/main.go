package main

import "fmt"

// EXPECTED OUTPUT
//  100
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	num1, _ := multiple()

	fmt.Println(num1)

	// THEN UNCOMMENT THE CODE BELOW

	// fmt.Println(b)
}

// multiple is a function that returns multiple int values
func multiple() (int, int) {
	return 100, 200
}