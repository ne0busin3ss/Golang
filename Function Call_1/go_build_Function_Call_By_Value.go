// Go program to illustrate the
// concept of the call by value
package main

import "fmt"

// function which swap values
func swap(x, y int) int {

	// taking a temporary variable
	var tmp int

	tmp = x
	x = y
	y = tmp

	return tmp
}

// Main function
func main() {

	var f int = 700
	var s int = 900

	fmt.Printf("Values Before Function Call\n")
	fmt.Printf("f = %d and s = %d\n", f, s)

	// call by values
	swap(f, s)

	fmt.Printf("\nValues After Function Call\n")
	fmt.Printf("f = %d and s = %d", f, s)
}
