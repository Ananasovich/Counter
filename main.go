package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filname := os.Args[1]
	data, err := ioutil.ReadFile(filname)
	if err != nil {
		fmt.Println("Ended with error", err)
		os.Exit(2)
	}

	fmt.Println("Number of lines in your file is", countLines(string(data)))
	fmt.Println("Number of runes in your file is", countRunes(string(data)))
}

func countLines(data string) int {
	return strings.Count(data, "\n") + 1
}

func countRunes(data string) int {
	return strings.Count(data, "") - 1
}
