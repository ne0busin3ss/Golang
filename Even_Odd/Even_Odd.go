//// Simple Program to Check Entered Number is Even or Odd
package main

import "fmt"

func main() {
	fmt.Print("Dude... enter a number please : ")
	var n int
	fmt.Scanln(&n)

	/* Conditional Statement if..., else...*/
	if n%2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is an odd number")
	}
}

/*Example to check Number is even or odd
In the above program. Scanlin function is used to take the input number from user console ie command line.
user entered number is stored in number variableNow checking number is even or odd, we used modulus
operator % and check remainder is zero with divisible by 2. conditional structure if else statement  is used.
the if the number is divisible by 2, then it is even number, an else odd number
*/
