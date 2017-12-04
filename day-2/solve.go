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
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) == 0 {
			continue
		}
		max, _ := strconv.Atoi(values[0])
		min, _ := strconv.Atoi(values[0])
		for i := range values {
			val, _ := strconv.Atoi(values[i])
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
		sum += max - min
	}
	fmt.Println(sum)
}
