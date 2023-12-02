package main

import (
	"fmt"
	"regexp"
)

func main() {
	const pattern = `((one)|(two)|(three)|(four)|
            (five)|(six)|(seven)|(eight)
            |(nine)|(zero)
            |(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9)|(0))`
	finalDigitArray := []string{}
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch("two1ninetwo", -1)
	for _, v := range matches {
		finalDigitArray = append(finalDigitArray, v[1])
	}
	fmt.Println(finalDigitArray)
}
