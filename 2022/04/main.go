package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 4 - Camp cleanup
func main() {
	// Read in file, strip out windows line endings
	input, _ := os.ReadFile("input.txt.txt")

	list := strings.Split(string(input), "\r\n")
	// Initialise counters for both tasks
	containedPairs := 0
	overlappingPairs := 0
	// Isolate the range values for each pair of elves and convert to integer
	for i := 0; i < len(list); i++ {
		elf1Start, _ := strconv.Atoi(strings.Split(strings.Split(list[i], ",")[0], "-")[0])
		elf1End, _ := strconv.Atoi(strings.Split(strings.Split(list[i], ",")[0], "-")[1])
		elf2Start, _ := strconv.Atoi(strings.Split(strings.Split(list[i], ",")[1], "-")[0])
		elf2End, _ := strconv.Atoi(strings.Split(strings.Split(list[i], ",")[1], "-")[1])

		// Task 1 - increment counter only if one elf's range completely enclose the other's
		if (elf1Start <= elf2Start && elf1End >= elf2End) || (elf1Start >= elf2Start && elf1End <= elf2End) {
			containedPairs += 1
		}
		// Task 2 - increment counter for any overlap
		if (elf1End >= elf2Start && elf1Start <= elf2End) || (elf2End >= elf1Start && elf2Start <= elf1End) {
			overlappingPairs += 1
		}
	}
	fmt.Println(containedPairs)   // Part 1 answer
	fmt.Println(overlappingPairs) // Part 2 answer
}
