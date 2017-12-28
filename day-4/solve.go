package main

import (
	"bufio"
	"fmt"
	"os"
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
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		set := make(map[string]bool)
		valid := true
		for _, word := range words {
			if set[word] {
				valid = false
				break
			}
			set[word] = true
		}
		if valid {
			sum += 1
		}
	}
	fmt.Println(sum)
}
