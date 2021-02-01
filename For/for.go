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
	for j := 7; j <= 10; j++ {
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
	for n := 0; n <= 11; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

/*Example to check Number is even or odd
In the above program. Scanlin function is used to take the input number from user console ie command line. user entered a number is stored in number variable

Now checking number is even or odd, we used modulus operator % and check remainder is zero with divisible by 2. conditional structure if else statement  is used. the if the number is divisible by 2, then it is even number, and else it is an odd number. The continue keyword tells Go to skip even numbers and display odd. 
*/