// Day 5 - Supply Stacks https://adventofcode.com/2022/day/5
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Read in the file
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(file), "\r\n")
	grid := input[:8]

	// Convert the grid to a map with each stack containing an array of its elements
	stacks := make(map[int][]string)
	// Generate from bottom up so items can be taken from end of slice
	for i := 7; i >= 0; i-- {
		for y := 1; y < ((len(grid[i])+1)/4)+1; y++ {
			if string(grid[i][(y*4-3)-1]) == "[" {
				item := string(grid[i][(y*4-3)-1 : ((y*4-3)-1)+3])
				stacks[y] = append(stacks[y], item)
			}
		}
	}
	stacksPart2 := make(map[int][]string)
	for k, v := range stacks {
		stacksPart2[k] = v
	}

	instructions := input[10:]

	for i := 0; i < len(instructions); i++ {
		instruction := strings.Split(instructions[i], " ")

		numToMove, _ := strconv.Atoi(instruction[1])
		numFrom, _ := strconv.Atoi(instruction[3])
		numTarget, _ := strconv.Atoi(instruction[5])

		// Part 1 Solution
		itemsToMove := stacks[numFrom][len(stacks[numFrom])-numToMove : len(stacks[numFrom])]
		stacks[numFrom] = stacks[numFrom][:len(stacks[numFrom])-numToMove]
		for item := range itemsToMove {
			stacks[numTarget] = append(stacks[numTarget], itemsToMove[numToMove-item-1])
		}

		// Part 2 Solution
		itemsToMove2 := stacksPart2[numFrom][len(stacksPart2[numFrom])-numToMove : len(stacksPart2[numFrom])]
		stacksPart2[numFrom] = stacksPart2[numFrom][:len(stacksPart2[numFrom])-numToMove]
		for _, item := range itemsToMove2 {
			stacksPart2[numTarget] = append(stacksPart2[numTarget], item)
		}

	}
	fmt.Println(stacks)      // Part 1 Answer
	fmt.Println(stacksPart2) // Part 2 Answer
}
