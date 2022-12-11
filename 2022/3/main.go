package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Read in file
	file, err := os.ReadFile("input.txt.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Split(string(file), "\n")
	priority := make(map[string]int)

	// Lowercase
	for i := 0; i < 26; i++ {
		str := string(rune('a' + i))
		priority[str] = i + 1
	}

	// Uppercase
	for i := 0; i < 26; i++ {
		str := string(rune('A' + i))
		priority[str] = i + 27
	}

	prioritySum := 0

	// Iterate through each rucksack
	for i := 0; i < len(rucksacks); i++ {
		// Split into the two compartments
		compartment1 := rucksacks[i][:(len(rucksacks[i]) / 2)]
		compartment2 := rucksacks[i][(len(rucksacks[i]) / 2):]
		for _, c := range string(compartment1) {
			if strings.Contains(string(compartment2), string(c)) {
				prioritySum += priority[string(c)]
				break
			}
		}
	}

	// Part 1 answer
	fmt.Printf("%v\n", prioritySum)

	// Part 2 solution
	prioritySum = 0
	for i := 1; i < (len(rucksacks)/3)+1; i++ {
		firstElf := rucksacks[(i*3-2)-1]
		secondElf := rucksacks[(i*3-1)-1]
		thirdElf := rucksacks[(i*3)-1]
		for _, c := range string(firstElf) {
			if strings.Contains(string(secondElf), string(c)) && strings.Contains(string(thirdElf), string(c)) {
				prioritySum += priority[string(c)]
				break
			}
		}
	}

	// Part 2 answer
	fmt.Print(prioritySum)

}
