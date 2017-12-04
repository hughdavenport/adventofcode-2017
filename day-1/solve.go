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
		fmt.Println(os.Args[0] + " <file>")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	captcha := strings.TrimSpace(string(data))
	sum := 0
	for i, c := range captcha[:len(captcha)-1] {
		if captcha[i+1] == captcha[i] {
			sum += int(c) - int('0')
		}
	}
	if captcha[0] == captcha[len(captcha)-1] {
		sum += int(captcha[0]) - int('0')
	}

	fmt.Println(sum)
}
