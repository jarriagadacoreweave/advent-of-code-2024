package main

import (
	"fmt"
  "aoc/utils"
  "strconv"
  "strings"
)

func main() {
	// Fetch data using utils function
	input, statusCode, err := utils.FetchData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the fetched data and status code
	fmt.Println("Status Code:", statusCode)
	//fmt.Println("Data:", input)
	// Initialize sum for valid results
	sum := 0
	n := len(input)
	i := 0

	// Iterate through the input string
	for i < n {
		// Look for the start of a "mul(" instruction
		if i+3 < n && input[i:i+4] == "mul(" {
			// Find the closing parenthesis for this mul block
			j := i + 4
			for j < n && input[j] != ')' {
				j++
			}

			// Ensure we found a closing parenthesis and validate the structure
			if j < n {
				// Extract the arguments inside "mul(...)"
				args := input[i+4 : j]
				// Check if arguments are valid
				if validateMul(args) {
					// Split and process arguments
          fmt.Println("Block:", args)
					parts := strings.Split(args, ",")
					x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
					y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
					sum += x * y
				}
			}

			// Move the index past the closing parenthesis
			i = j + 1
		} else {
			// Move to the next character
			i++
		}
	}

	// Output the result
	fmt.Println("Sum of all valid mul instructions:", sum)
}

// validateMul checks if the string inside "mul(...)" is valid
func validateMul(content string) bool {
	parts := strings.Split(content, ",")
	if len(parts) != 2 {
		return false
	}

	// Check if both parts are integers
	_, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	_, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	return err1 == nil && err2 == nil
}

