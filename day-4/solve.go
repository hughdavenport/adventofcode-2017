package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type runeArray []rune

func (s runeArray) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runeArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runeArray) Len() int {
	return len(s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println()
		fmt.Println(os.Args[0] + " <file> [part2]")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	defer file.Close()
	sortChars := false
	if len(os.Args) == 3 {
		sortChars = true
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		set := make(map[string]bool)
		valid := true
		for _, word := range words {
			if sortChars {
				runes := []rune(word)
				sort.Sort(runeArray(runes))
				word = string(runes)
			}
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
