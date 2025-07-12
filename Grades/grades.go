package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// passingGrade defines the minimum grade to pass.
const passingGrade = 70.0

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
		log.Fatal("Failed to read input: ", err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal("Invalid input. Please enter a number: ", err)
	}

	status := getStatus(grade)
	fmt.Println("A grade of", grade, "is", status)
}
