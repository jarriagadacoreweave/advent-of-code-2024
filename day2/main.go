package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
  "strconv"
	"strings"
)

// isSafe checks if a report is safe according to the given rules.
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false // A report with fewer than two levels cannot be safe.
	}

	// Determine if the report is increasing or decreasing.
	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check the difference is within the range 1 to 3 or -1 to -3.
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Check consistency of direction (all increasing or all decreasing).
		if (diff > 0) != isIncreasing {
			return false
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true // If already safe, no need to apply the dampener.
	}

	// Try removing each level and check if the resulting report is safe.
	for i := 0; i < len(report); i++ {
		// Create a new report without the current level.
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		// Check if the modified report is safe.
		if isSafe(newReport) {
			return true
		}
	}

	return false
}

func processRawData(rawData string) [][]int {
	lines := strings.Split(strings.TrimSpace(rawData), "\n")
	var reports [][]int

	for _, line := range lines {
		numbers := strings.Fields(line)
		var report []int
		for _, num := range numbers {
			value, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Invalid number '%s', skipping report\n", num)
				report = nil // Skip this line entirely if there's an invalid number.
				break
			}
			report = append(report, value)
		}
		if report != nil {
			reports = append(reports, report)
		}
	}

	return reports
}

func main() {
	url := os.Getenv("AOC_URL")
	cookieName := os.Getenv("COOKIE_NAME")
	cookieValue := os.Getenv("COOKIE_VALUE")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Status Code:", resp.StatusCode)
  //fmt.Println("Data:")
  //fmt.Println(string(rawData))
  reports := processRawData(string(rawData))

	safeCount := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			//fmt.Printf("Report %d: Safe\n", i+1)
			safeCount++
		} else {
			//fmt.Printf("Report %d: Unsafe\n", i+1)
		}
	}

	fmt.Printf("Total safe reports with Problem Dampener: %d\n", safeCount)
}

