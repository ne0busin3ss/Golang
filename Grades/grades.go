package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// passingGrade defines the minimum grade to pass.
const passingGrade = 75.0

// getStatus determines if a grade is passing or failing.
func getStatus(grade float64) string {
	if grade >= passingGrade {
		return "passing"
	}
	return "failing"
}

func main() {
	fmt.Print("Please enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input: %v\n", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input. Please enter a number: %v\n", err)
		os.Exit(1)
	}

	if grade < 0 || grade > 100 {
		fmt.Fprintln(os.Stderr, "Grade must be between 0 and 100.")
		os.Exit(1)
	}

	status := getStatus(grade)
	fmt.Println("A grade of", grade, "is", status)
}
