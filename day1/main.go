package main

import (
  "fmt"
  "io"
  "os"
  "net/http"
	"strconv"
	"strings"
  "sort"
  "math"
)

func min(slice []int) (int, int, error) {
	if len(slice) == 0 {
		return 0, -1, fmt.Errorf("slice is empty")
	}

	smallest := slice[0]
	index := 0

	for i, num := range slice {
		if num < smallest {
			smallest = num
			index = i
		}
	}

	return smallest, index, nil
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

  // Split the raw data into lines
	lines := strings.Split(string(rawData), "\n")

	// Initialize slices for left and right columns
	var left []int
	var right []int

	// Iterate through each line
	for _, line := range lines {
		// Split the line into two parts by spaces
		columns := strings.Fields(line)

		// Parse the left and right columns into integers
		if len(columns) == 2 { // Ensure there are exactly two columns
			leftVal, err1 := strconv.Atoi(columns[0])
			rightVal, err2 := strconv.Atoi(columns[1])
			if err1 == nil && err2 == nil {
				left = append(left, leftVal)
				right = append(right, rightVal)
			} else {
				fmt.Println("Error parsing integers:", err1, err2)
			}
		}
	}
	// Ensure both lists have the same length
	if len(left) != len(right) {
		fmt.Println("Error: Lists are not the same length.")
		return
	}

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Compute total distance
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		totalDistance += distance
	}

	// Print the result
	fmt.Println("Total Distance:", totalDistance)

	// Create a frequency map for the right list
	frequencyMap := make(map[int]int)
	for _, num := range right {
		frequencyMap[num]++
	}

	// Calculate the total similarity score
	totalSimilarityScore := 0
	for _, num := range left {
		frequency := frequencyMap[num]
		totalSimilarityScore += num * frequency
	}

	// Print the result
	fmt.Println("Total Similarity Score:", totalSimilarityScore)

	// Print the resulting slices
	// fmt.Println("Left Column:", left)
	// fmt.Println("Right Column:", right)

}
