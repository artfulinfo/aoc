// Day 6 - Tuning Trouble
package main

import (
	"fmt"
	"os"
	"strings"
)

// Check for unique character combinations
func checkUnique(p []byte) bool {
	unique := true
	for _, char := range p {
		count := strings.Count(string(p), string(char))
		if count > 1 {
			unique = false
			break
		}
	}
	return unique
}

// Locate the first instance of a successful start packet, returning the packet and location
func locateSignal(u int) string {
	input, _ := os.ReadFile("input.txt")
	var successPacket string
	for i := u; i < len(input); i++ {
		startPacket := input[i-u : i]
		if checkUnique(startPacket) == true {
			successPacket = fmt.Sprintf("%v %s", i, startPacket)
			break
		}
	}
	return successPacket
}

func main() {
	// Part 1 answer
	fmt.Println(locateSignal(4))

	// Part 2 answer
	fmt.Println(locateSignal(14))
}
