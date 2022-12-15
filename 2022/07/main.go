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

type Directory struct {
	name     string
	parent   *Directory
	contents map[string]*Directory
	files    []File
}

func main() {
	input, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(input), "\r\n")
	// Create the root directory
	root := &Directory{name: "/"}
	cwd := root
	for i := 0; i < len(commands); i++ {
		_, err := strconv.Atoi(commands[i][:1])
		c := strings.Split(commands[i], " ")
		if err == nil {
			s, _ := strconv.Atoi(c[0])
			cwd.files = append(cwd.files, File{c[1], s})
			//tree[i][500] = File{c[1], s}
		} else if commands[i][:4] == "$ cd" && c[2][:1] != "." {
			subdir := &Directory{name: commands[i][5:]}
			contents := make(map[string]*Directory)
			contents[subdir.name] = &Directory{name: subdir.name, parent: cwd}
			cwd.contents = contents
			cwd = cwd.contents[subdir.name]
			fmt.Printf("%v\n", cwd)
		}

	}
	//fmt.Println(root)
}
