package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(input), "\r\n")
	tree := make(map[int][]File)
	for i := 0; i < len(commands); i++ {
		_, err := strconv.Atoi(commands[i][:1])
		if err == nil {
			c := strings.Split(commands[i], " ")
			s, _ := strconv.Atoi(c[0])
			tree[i] = append(tree[i], File{c[1], s})
			tree[i][500] = File{c[1], s}
			fmt.Println(tree[i][0])
		}

	}
	fmt.Println(tree)
}
