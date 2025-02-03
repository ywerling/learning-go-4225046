// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	// "strings"
)

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the the result of adding 2 numbers 
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
    // Your code goes here.

	// Convert the first string to a float64
	string1, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		// error handling
    fmt.Printf("error")
	}
	float1 := float64(string1)
	// Convert the second string to a float64
	string2, err := strconv.ParseFloat(value2, 64)
	if err != nil {
    // error handling
		fmt.Printf("error")
	}
	float2 := float64(string2)
	// Calculate and return the result

	return float1 + float2
}
