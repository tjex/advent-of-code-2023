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
	stringInt := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	scanner := bufio.NewScanner(file)

	var answer int
	var aChecked, bChecked int

	const pattern = `((one)|(two)|(three)|(four)|
            (five)|(six)|(seven)|(eight)
            |(nine)|(zero)
            |(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9)|(0))`
	digitAsStringArray := []string{}
	r := regexp.MustCompile(pattern)

	for scanner.Scan() {
		digitAsStringArray = nil
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		for _, v := range matches {
			digitAsStringArray = append(digitAsStringArray, v[1])
		}
		// fmt.Println()
		a := strings.TrimSpace(digitAsStringArray[0])
		b := strings.TrimSpace(digitAsStringArray[len(digitAsStringArray)-1])

		firstNumeral, err := strconv.Atoi(a)
		aChecked = firstNumeral
		if err != nil {
			aChecked = stringInt[a]
		}

		lastNumeral, err := strconv.Atoi(b)
		bChecked = lastNumeral
		if err != nil {
			bChecked = stringInt[b]
		}

		if len(digitAsStringArray) > 1 {
			// fmt.Println(aChecked, bChecked)
			doubleDig := aChecked*10 + bChecked
			answer += doubleDig
		} else {
			fmt.Println(scanner.Text())
			fmt.Println(aChecked, bChecked)
			answer += aChecked
		}

		// fmt.Println(reflect.TypeOf(digitAsStringArray[0]),  reflect.TypeOf(digitAsStringArray[len(digitAsStringArray)-1]))
		// fmt.Println(scanner.Text())
		// fmt.Println("length", len(digitAsStringArray))
		// fmt.Println(aChecked, bChecked)
		// fmt.Printf("first as string %q\n", a)
		// fmt.Printf("second as string %q\n", b)
		// fmt.Println("dig", aChecked, bChecked)
		// fmt.Println(doubleDig)
		//       fmt.Println("running total", answer)
		// fmt.Println()
	}
	fmt.Println(answer)
}

func main() {
	// runeSol()
	regexSol()

}
