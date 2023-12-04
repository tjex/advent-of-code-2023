package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// input_file, err := os.Open("sample.txt")
	// input_file, err := os.Open("sample2.txt")
	input_file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input_file.Close()
	sc := bufio.NewScanner(input_file)

	var theInput []string

	for sc.Scan() {
		theInput = append(theInput, sc.Text())
	}

	partOne(theInput)
	partTwo(theInput)
}

func partOne(theInput []string) {

	totalSum := 0

	for _, line := range theInput {

		var currentLine []int

		for _, i := range line {
			// If the rune is a number, then convert to an int and add to slice
			if isNumber(string(i)) {
				num, _ := strconv.Atoi(string(i))
				currentLine = append(currentLine, num)
			}
		}

		// Grabs and formats 1st and last ints in line to a double digit number
		firstDigit := currentLine[0]
		lastDigit := currentLine[len(currentLine)-1]

		doubleDigit := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		num, _ := strconv.Atoi(doubleDigit)
		totalSum += num
	}
	fmt.Printf("Part 1 answer is: %v\n", totalSum)
}

func partTwo(theInput []string) {

	totalSum := 0
	// Regex to parse out both the spelled out numbers and digits
	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for _, line := range theInput {

		var currentLine []string

		for i := range line {
			found := re.FindString(line[i:])
			// Don't want to append empty strings to our slice
			if found != "" {
				currentLine = append(currentLine, found)
			}
		}
        // fmt.Println(currentLine)

		// Converts "one" to "1", etc.
		for index, word := range currentLine {
			if !isNumber(word) && word != "" {
				currentLine[index] = wordToNum(word)
			}
		}

		// Grabs and formats 1st and last strings in line to a double digit number
		firstDigit := currentLine[0]
		lastDigit := currentLine[len(currentLine)-1]

		doubleDigit := fmt.Sprintf("%s%s", firstDigit, lastDigit)
		num, _ := strconv.Atoi(doubleDigit)
		totalSum += num
	}
	fmt.Printf("Part 2 answer is: %d\n", totalSum)
}

// Check if the string is a number by returning whether or not there was
// an error during type conversion
func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

// Converts spelled out numbers to string of digit
// e.g. "one" becomes "1"
func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return ""
	}
}

