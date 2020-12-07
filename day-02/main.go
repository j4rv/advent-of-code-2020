package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const policyRgx = "(\\d+)-(\\d+) (\\w): (.*)"
const minGroupIndex, maxGroupIndex, charGroupIndex, passGroupIndex = 1, 2, 3, 4

func main() {
	r := regexp.MustCompile(policyRgx)

	// Part One
	var counter1 int
	for _, rowDataStr := range strings.Split(data, "\n") {
		rowData := r.FindStringSubmatch(rowDataStr)
		min := mustParseInt(rowData[minGroupIndex])
		max := mustParseInt(rowData[maxGroupIndex])
		char := rowData[charGroupIndex]
		pass := rowData[passGroupIndex]

		// minus one because even if the string does not contain the char, its length is still one
		charCount := len(strings.Split(pass, char)) - 1

		// policy: char must appear at least "min" times and "max" times at most
		if min <= charCount && charCount <= max {
			//log.Println(rowDataStr, charCount, "valid!")
			counter1++
		}
	}
	log.Println("Part One solution:", counter1)

	// Part Two
	var counter2 int
	for _, rowDataStr := range strings.Split(data, "\n") {
		rowData := r.FindStringSubmatch(rowDataStr)

		// substract one to make them "start at zero"
		index1 := mustParseInt(rowData[minGroupIndex]) - 1
		index2 := mustParseInt(rowData[maxGroupIndex]) - 1
		char := rowData[charGroupIndex]
		pass := rowData[passGroupIndex]

		// policy: char must appear exactly once at either index1 or index2
		charCount := 0
		if string(pass[index1]) == char {
			charCount++
		}
		if string(pass[index2]) == char {
			charCount++
		}
		if charCount == 1 {
			//log.Println(rowDataStr, charCount, "valid!")
			counter2++
		}
	}
	log.Println("Part Two solution:", counter2)
}

func mustParseInt(str string) int {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(res)
}
