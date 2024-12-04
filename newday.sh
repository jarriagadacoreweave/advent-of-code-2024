#!/bin/bash

#create a new day

day=$1
dir="day${day}"

echo -e "Generating day $day..."

mkdir day${day}

cat > ./${dir}/main.go << EOF
package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	input, statusCode, err := utils.FetchData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status Code:", statusCode)
	fmt.Println("Data:", input)
}
EOF

