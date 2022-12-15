package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	heights := string(input)
	rows := strings.Split(heights, "\r\n")
	// Part 1 solution
	visibleTrees := (len(rows)+len(rows[0]))*2 - 4 // Perimeter trees as they are all visible
	for i := 1; i < len(rows)-1; i++ {
		for t := 1; t < len(rows[i])-1; t++ {
			visibleLeft := true
			for _, oth := range rows[i][0:t] {
				if int(oth) >= int(rows[i][t]) {
					visibleLeft = false
				}
			}
			visibleRight := true
			for _, oth := range rows[i][t+1 : len(rows[i])] {
				if int(oth) >= int(rows[i][t]) {
					visibleRight = false
				}
			}
			visibleTop := true
			for _, oth := range rows[0:i] {
				if int(oth[t]) >= int(rows[i][t]) {
					visibleTop = false
				}
			}
			visibleBottom := true
			for _, oth := range rows[i+1:] {
				if int(oth[t]) >= int(rows[i][t]) {
					visibleBottom = false
				}
			}
			// Count the tree as visible if it can be seen from any direction
			if visibleLeft == true || visibleRight == true || visibleTop == true || visibleBottom == true {
				visibleTrees += 1
			}
		}

	}
	// Part 2 solution
	scenicRecord := 0
	for i := 1; i < len(rows)-1; i++ {
		for t := 1; t < len(rows[i])-1; t++ {
			scenicLeft := 0
			for o := t - 1; o >= 0; o-- {
				if int(rows[i][o]) < int(rows[i][t]) {
					scenicLeft += 1
				} else if int(rows[i][o]) >= int(rows[i][t]) {
					scenicLeft += 1
					break
				}
			}
			scenicRight := 0
			for o := t + 1; o < len(rows[i]); o++ {
				if int(rows[i][o]) < int(rows[i][t]) {
					scenicRight += 1
				} else if int(rows[i][o]) >= int(rows[i][t]) {
					scenicRight += 1
					break
				}
			}
			scenicTop := 0
			for o := i - 1; o >= 0; o-- {
				if int(rows[o][t]) < int(rows[i][t]) {
					scenicTop += 1
				} else if int(rows[o][t]) >= int(rows[i][t]) {
					scenicTop += 1
					break
				}
			}
			scenicBottom := 0
			for o := i + 1; o < len(rows); o++ {
				if int(rows[o][t]) < int(rows[i][t]) {
					scenicBottom += 1
				} else if int(rows[o][t]) >= int(rows[i][t]) {
					scenicBottom += 1
					break
				}
			}
			scenicScore := scenicLeft * scenicRight * scenicTop * scenicBottom
			if scenicScore > scenicRecord {
				scenicRecord = scenicScore
			}

		}
	}
	// Part 1 answer
	fmt.Println(visibleTrees)
	// Part 2 answer
	fmt.Println(scenicRecord)
}
