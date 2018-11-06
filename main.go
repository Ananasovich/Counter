package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ended with error while read file", err)
		os.Exit(2)
	}
	newfilename := "counted" + filename
	err = ioutil.WriteFile(newfilename, result(data, filename), 0777)
	if err != nil {
		fmt.Println("Ended with error while write result to file", err)
		os.Exit(2)
	}

}

func countLines(data string) int {
	return strings.Count(data, "\n") + 1
}

func countRunes(data string) int {
	return strings.Count(data, "") - 1
}

func countWords(data string) int {
	return len(strings.Fields(data))
}

func result(data []byte, filename string) []byte {
	dataS := string(data)
	dataToWrite := "Counted for file " + filename + "\nNumber of lines in this file is " +
		strconv.Itoa(countLines(dataS)) + ".\nNumber of words in this file is " +
		strconv.Itoa(countWords(dataS)) + ".\nNumber of runes in this file is " +
		strconv.Itoa(countRunes(dataS)) + "."
	return []byte(dataToWrite)
}
