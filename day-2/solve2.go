package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println()
		fmt.Println(os.Args[0] + " <file> [part2]")
		os.Exit(1)
	}
	filename := os.Args[1]
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	scanner := bufio.NewScanner(bufio.NewReader(fd))
	sum := 0
Loop:
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) == 0 {
			continue
		}
		for i := range values {
			for j := range values {
				if i == j {
					continue
				}
				val1, _ := strconv.Atoi(values[i])
				val2, _ := strconv.Atoi(values[j])
				if val1%val2 == 0 {
					sum += val1 / val2
					continue Loop
				}
			}
		}
	}
	fmt.Println(sum)
}
