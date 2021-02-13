// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 15; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 13; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

/*'n' is the vaiable for the number. %2 == 0 tells go to divide the number by 2 and if the remander is 0, it is even; continue to the next iteration
-don't display even numbers.*/
