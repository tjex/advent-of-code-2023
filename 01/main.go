package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var answer int
	for scanner.Scan() {
		r := []rune(scanner.Text())
		wordRange := len(r)
		digitArray := []rune{}
		for i, rune := range r {
			if unicode.IsNumber(rune) && i != wordRange {
				digitArray = append(digitArray, rune)
			}
			if i+1 == wordRange {
				firstDigit := digitArray[0]
				secondDigit := digitArray[len(digitArray)-1]
				digitAsArray := string(firstDigit) + string(secondDigit)
				digitAsInt, _ := strconv.Atoi(digitAsArray)

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
