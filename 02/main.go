package main

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func runeSol() {
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	digitArray := []rune{}
	var answer int
	// s := []string{}

	for scanner.Scan() {
		r := []rune(scanner.Text())
		wordRange := len(r)
		for i, rune := range r {
			if unicode.IsNumber(rune) && i != wordRange {
				digitArray = append(digitArray, rune)
			}
			if i+1 == wordRange {
				firstDigit := digitArray[0]
				secondDigit := digitArray[len(digitArray)-1]
				digitAsString := string(firstDigit) + string(secondDigit)
				digitAsInt, _ := strconv.Atoi(digitAsString)

				// fmt.Println()
				// fmt.Println(scanner.Text())
				// fmt.Println("rune array conv to string:", string(digitArray))
				// fmt.Println("int derrived:", digitAsInt)
				answer += digitAsInt
				digitArray = nil
			}

		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(answer)
}

func regexSol() {
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}
	intNameToString := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	scanner := bufio.NewScanner(file)

	var answer int

	const pattern = `^(\d|one|two|three|four|five|six|seven|eight|nine)`
	r := regexp.MustCompile(pattern)
	currentLine := []string{}
	inputArray := []string{}

	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}
	for _, line := range inputArray {
		currentLine = nil
		for i := range line {
			match := r.FindString(line[i:])
			if match != "" {
				currentLine = append(currentLine, match)
			}

			// fmt.Println(currentLine)
		}

		for i, word := range currentLine {
			_, err := strconv.Atoi(word)
			if err != nil {
				currentLine[i] = intNameToString[word]
			}
		}
		a := strings.TrimSpace(currentLine[0])
		b := strings.TrimSpace(currentLine[len(currentLine)-1])

		doubleDig := fmt.Sprintf("%s%s", a, b)

		num, _ := strconv.Atoi(doubleDig)
		// fmt.Println("digit:", num)
		// fmt.Println()
		answer += num
	}

	// fmt.Println(reflect.TypeOf(currentLine[0]),  reflect.TypeOf(currentLine[len(currentLine)-1]))
	// fmt.Println(scanner.Text())
	// fmt.Println("length", len(currentLine))
	// fmt.Println(aChecked, bChecked)
	// fmt.Printf("first as string %q\n", a)
	// fmt.Printf("second as string %q\n", b)
	// fmt.Println(doubleDig)
	//       fmt.Println("running total", answer)
	// fmt.Println()
	fmt.Println(answer)
}

func main() {
	// runeSol()
	regexSol()

}
