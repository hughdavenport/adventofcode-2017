package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println()
		fmt.Println(os.Args[0] + " <file>")
		os.Exit(1)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	instructions := make([]int, len(lines))
	for i, _ := range lines {
		instructions[i], err = strconv.Atoi(lines[i])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}
	}

	pc := 0
	sum := 0
	for {
		instructions[pc] += 1
		pc += (instructions[pc] - 1)
		sum += 1
		if pc < 0 || pc >= len(instructions) {
			break
		}
	}

	fmt.Println(sum)
}
