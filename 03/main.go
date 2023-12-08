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
var pTwoAnswer int

func main() {
	inputData := extractInputData()
	symData := extractSymbolIndicies(inputData)
	numData := extractNumberIndicies(inputData)
	pOneAnswer := returnSumOfValidNumbers(inputData, symData, numData)
    gearRatios := findGearRatios(inputData, symData, numData)
    for _, ratio := range gearRatios {
        pTwoAnswer += ratio
    }

	fmt.Println("part one:", pOneAnswer)
	fmt.Println("part two:", pTwoAnswer)

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

func printDataAtRange(row, rangeLeft, rangeRight int) {

	inputData := extractInputData()
	for i, line := range inputData {
		if i == row-1 || i == row || i == row+1 {
			fmt.Println(line[rangeLeft-1 : rangeRight+1])
		}
	}
}

func returnNumberAtRange(row, rangeLeft, rangeRight int) int {
	inputData := extractInputData()
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

func doesRangeContainGear(row, rangeLeft, rangeRight int, inputData []string) bool {
	for i, line := range inputData {
		if i == row {
			// get the text at the given range
			s := line[rangeLeft:rangeRight]
			matched, err := regexp.MatchString(`\*`, s)
			if err != nil {
				fmt.Println(err)
			}
			return matched
		}

	}
	return false
}

func doesSymbolSurroundNumber(rowToCheck, rangeLeft, rangeRight int, symData [][][]int) bool {
	// rangeRight is not extended by +1 because it is already
	// the cell "to the right of the number" (due to an array's end range value being "up to but not
	// including").
	cellLeftofNumber := rangeLeft - 1
	cellRightofNumber := rangeRight
	for i, line := range symData {
		// check if symbol exists above number (including diagonals)
		if i >= rowToCheck-1 || i <= rowToCheck+1 {
			for _, row := range line {
				if row[0] >= cellLeftofNumber && row[0] <= cellRightofNumber {
					return true
					// fmt.Println("symbol above and/or below number")
				}
			}
		}
	}
	return false
}

func doNumbersSurroundGear(rowWithGear, gearCell int, numData [][][]int) (int, int) {
	var numTop int
	var numBot int
	for i, line := range numData {
		// only check lines above and below. Numbers next to gears are not valid
		if i == rowWithGear-1 {

			for _, row := range line {
				leftBorder := row[0] - 1
				rightBorder := row[1]
				if gearCell >= leftBorder && gearCell <= rightBorder {
					// fmt.Println(gearCell, line)
					// fmt.Println("gear within number range")
					// printDataAtRange(i, row[0], row[1])
					numTop = returnNumberAtRange(i, row[0], row[1])
				}
			}

		}
		if i == rowWithGear+1 {
			for _, row := range line {
				leftBorder := row[0] - 1
				rightBorder := row[1]
				if gearCell >= leftBorder && gearCell <= rightBorder {
					// fmt.Println(gearCell, line)
					// fmt.Println("gear within number range")
					// printDataAtRange(i, row[0], row[1])
					numBot = returnNumberAtRange(i, row[0], row[1])
				}
			}
		}
	}

    return numTop, numBot
}

// first find a gear, and then check for numbers around it
func findGearRatios(inputData []string, symData, numData [][][]int) []int {
    var gearRatioArray []int
	for i, line := range symData {
		for _, row := range line {
			rangeLeft := row[0]
			rangeRight := row[1]
			if doesRangeContainGear(i, rangeLeft, rangeRight, inputData) {
                numTop, numBot := doNumbersSurroundGear(i, rangeLeft, numData)
                gearRatio := numTop * numBot
                gearRatioArray = append(gearRatioArray, gearRatio)
			}
		}
	}
    return gearRatioArray

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
			if doesSymbolSurroundNumber(i, rangeLeft, rangeRight, symData) {
				validNumber := returnNumberAtRange(i, rangeLeft, rangeRight)
				if validNumber != 0 {
					sum += validNumber
				}

			}
		}
	}
	return sum
}
