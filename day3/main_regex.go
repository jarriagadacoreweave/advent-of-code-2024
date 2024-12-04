package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input, statusCode, err := utils.FetchData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Code:", statusCode)
	// Define regex patterns
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	// Variables to track state and sum
	isEnabled := true
	sum := 0

	// Iterate over the input string to process instructions
	for len(input) > 0 {
		// Find the earliest match for mul(), do(), or don't()
		mulMatch := mulRegex.FindStringSubmatchIndex(input)
		doMatch := doRegex.FindStringIndex(input)
		dontMatch := dontRegex.FindStringIndex(input)

		// Determine the next instruction
		nextInstr := -1
		if mulMatch != nil && (doMatch == nil || mulMatch[0] < doMatch[0]) && (dontMatch == nil || mulMatch[0] < dontMatch[0]) {
			nextInstr = mulMatch[0]
			if isEnabled {
				// Process the mul() instruction
				x, _ := strconv.Atoi(input[mulMatch[2]:mulMatch[3]])
				y, _ := strconv.Atoi(input[mulMatch[4]:mulMatch[5]])
				sum += x * y
			}
			// Remove processed mul() instruction
			input = input[mulMatch[1]:]
		} else if doMatch != nil && (dontMatch == nil || doMatch[0] < dontMatch[0]) {
			nextInstr = doMatch[0]
			// Process the do() instruction
			isEnabled = true
			// Remove processed do() instruction
			input = input[doMatch[1]:]
		} else if dontMatch != nil {
			nextInstr = dontMatch[0]
			// Process the don't() instruction
			isEnabled = false
			// Remove processed don't() instruction
			input = input[dontMatch[1]:]
		} else {
			// No more instructions to process
			break
		}

		// If no valid instruction found, break
		if nextInstr == -1 {
			break
		}
	}

	// Output the result
	fmt.Println("Sum of all enabled mul instructions:", sum)
}
