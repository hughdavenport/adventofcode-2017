package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	captcha := strings.TrimSpace(string(data))
	gap := 1
	if len(os.Args) >= 3 {
		gap = len(captcha) / 2
	}
	sum := 0
	for i, c := range captcha {
		if captcha[(i+gap)%len(captcha)] == captcha[i] {
			sum += int(c) - int('0')
		}
	}

	fmt.Println(sum)
}
