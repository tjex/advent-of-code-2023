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
	checkNumberBoundaries(inputData, symbolIndicies, numberIndicies)
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

func extractNumberIndicies(inputData []string) [][][]int {
	var numbers [][][]int
	symbol := regexp.MustCompile(numberPat)
	for _, line := range inputData {
		found := symbol.FindAllStringIndex(line, -1)
		numbers = append(numbers, found)
	}
	numbers = numbers[:]

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

func convertTo2DArray(a [][][]int) [][]int {
	var twoDarray [][]int
	for _, line := range a {
		for _, row := range line {
			twoDarray = append(twoDarray, row)
		}
	}

	return twoDarray

}

func printDataAtRange(row, rangeLeft, rangeRight int, inputData []string) {
	for i, line := range inputData {
		if i == row {
			fmt.Println(line[rangeLeft:rangeRight])
		}

	}

}

// check if symbol exists in any neighbouring cell to a number, and if so return the numer
func checkNumberBoundaries(inputData []string, symData, numData [][][]int) {
	// dataRows := len(numData)
	// data arrays with different nRows will break this function
	if (len(symData) - len(numData)) != 0 {
		panic(`symbol and number indicies array length mismatch,
        something has gone wrong when extracting indicies for each type`)
	}

	// for each nIndicie group, check if:
	// - a symbol exists (+/- 1 cell in the same row)
	// - a symbol exists (within the same range OR +/- 1 cell in rows +/- 1)

	for i, line := range numData {
		for _, row := range line {
			rangeLeft := row[0]
			rangeRight := row[1]
			fmt.Println(rangeLeft, rangeRight)
			printDataAtRange(i, rangeLeft, rangeRight, inputData)
		}
	}

}
