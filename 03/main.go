package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const numberPat = `(\d+)`
const symbolPat = `[^a-zA-Z0-9.]`

func main() {
	inputData := extractInputData()
	symbolIndicies := extractSymbolIndicies(inputData)
	numberIndicies := extractNumberIndicies(inputData)
    for i := 0 ; i < len(numberIndicies); i++ {
		fmt.Println("number:", i, numberIndicies[i])
        fmt.Println("symbol", i, symbolIndicies[i])
        fmt.Println()
	}
}
func extractNumberIndicies(inputData []string) [][][]int {
	var numbers [][][]int
	symbol := regexp.MustCompile(numberPat)
	for _, line := range inputData {
		found := symbol.FindAllStringIndex(line, -1)
		numbers = append(numbers, found)
	}

	return numbers

}

func extractSymbolIndicies(inputData []string) [][][]int {
	var symbols [][][]int
	symbol := regexp.MustCompile(symbolPat)
	for _, line := range inputData {
		found := symbol.FindAllStringIndex(line, -1)
		symbols = append(symbols, found)
	}

	return symbols

}

func extractInputData() []string {
	var inputDataAsStringArray []string
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("file open error:", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		inputDataAsStringArray = append(inputDataAsStringArray, sc.Text())
	}

	return inputDataAsStringArray
}
