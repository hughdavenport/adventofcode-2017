package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println()
		fmt.Println(os.Args[0] + " <input>")
		os.Exit(1)
	}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	if input == 1 {
		fmt.Println(0)
		return
	}
	ringEnd := 1
	ringIndex := 1
	for ringEnd < input {
		ringIndex += 1
		ringEnd += 4 * (ringIndex - 1) * 2
	}
	ringStart := ringEnd - 4 * (ringIndex  - 1) * 2 + 1
	side := (ringEnd - ringStart + 1) / 4 + 1
	input -= ringStart - 1
	input %= side - 1
	mid := side / 2

	distance := ringIndex - 1 + int(math.Abs(float64(input - mid)))
	fmt.Println(distance)
}
