package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const numberPat = `(\d+)`
const symbolPat = `[^a-zA-Z0-9.]`

func main() {
	inputData := extractInputData()
	symbolIndicies := extractSymbolIndicies(inputData)
	numberIndicies := extractNumberIndicies(inputData)
	answer := returnSumOfValidNumbers(inputData, symbolIndicies, numberIndicies)
	fmt.Println(answer)
}
func extractInputData() []string {
	var inputDataAsStringArray []string
	file, err := os.Open("input-test.txt")
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

func printDataAtRange(row, rangeLeft, rangeRight int) {

	inputData := extractInputData()
	for i, line := range inputData {
		if i == row-1 || i == row || i == row+1 {
			fmt.Println(line[rangeLeft-1 : rangeRight+1])
		}
	}
	fmt.Println("----------------")
}

func returnNumberAtRange(row, rangeLeft, rangeRight int, inputData []string) int {
	for i, line := range inputData {
		if i == row {
			// get the text at the given range
			s := line[rangeLeft:rangeRight]
			validNumber, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("in returnNumberAtRange():", err)
			}
			return validNumber
		}

	}
	return 0

}
func doesSymbolSurroundNumberChat(rowToCheck, rangeLeft, rangeRight int, symData [][][]int) bool {
	cellLeftOfNumber := rangeLeft - 1
	cellRightOfNumber := rangeRight + 1

	for i, line := range symData {
		if i == rowToCheck {
			// Check if symbol exists before or after the number
			for _, row := range line {
				if row[0] == cellLeftOfNumber || row[0] == cellRightOfNumber {
					return true
				}
			}
		}

		// Check if symbol exists directly above or below the number
		if i == rowToCheck-1 || i == rowToCheck+1 {
			for _, row := range line {
				if row[0] >= cellLeftOfNumber && row[0] <= cellRightOfNumber {
					return true
				}
			}
		}
	}

	return false
}

func doesSymbolSurroundNumber(rowToCheck, rangeLeft, rangeRight int, symData [][][]int) bool {
	doesSurround := false
	cellLeftofNumber := rangeLeft - 1
	cellRightofNumber := rangeRight + 1
	for i, line := range symData {
		// check if symbol exists above number (including diagonals)
		if i >= rowToCheck-1 || i <= rowToCheck+1 {
			for _, row := range line {
				if row[0] >= cellLeftofNumber && row[0] <= cellRightofNumber {
					doesSurround = true
					// fmt.Println("symbol above and/or below number")
				}
			}
		}
	}
	return doesSurround
}

// check if symbol exists in any neighbouring cell to a number, and if so return the numer
func returnSumOfValidNumbers(inputData []string, symData, numData [][][]int) int {
	var sum int
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
			if doesSymbolSurroundNumberChat(i, rangeLeft, rangeRight, symData) {
				printDataAtRange(i, rangeLeft, rangeRight)
				validNumber := returnNumberAtRange(i, rangeLeft, rangeRight, inputData)
				if validNumber != 0 {
					sum += validNumber
				}

			}
		}
	}
	return sum
}
