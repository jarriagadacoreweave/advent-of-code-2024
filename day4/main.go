package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var directions = [][2]int{
	{0, 1},   // Horizontal right
	{1, 0},   // Vertical down
	{0, -1},  // Horizontal left
	{-1, 0},  // Vertical up
	{1, 1},   // Diagonal down-right
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
}

func CountOccurrences(grid []string, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	// grid movement
	checkWord := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+i*dx, y+i*dy
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != word[i] {
				return false
			}
		}
		return true
	}

	// move through every cell in the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				if checkWord(x, y, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	input, statusCode, err := utils.FetchData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Code:", statusCode)
	//fmt.Println("Data:", input)

	// Parse the input into a grid
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var grid []string
	for _, line := range lines {
		grid = append(grid, strings.TrimSpace(line))
	}

	word := "XMAS"
	count := CountOccurrences(grid, word)
	fmt.Println("XMAS Count:", count)
}
