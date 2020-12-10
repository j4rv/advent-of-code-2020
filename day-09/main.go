package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{}
	for _, dataLine := range strings.Split(data, "\n") {
		num, err := strconv.Atoi(dataLine)
		if err != nil {
			panic("not a number: " + dataLine)
		}
		numbers = append(numbers, num)
	}

	// Part One
	var invalidNum int
	for i := range numbers {
		if !isValid(i, numbers, 25) {
			invalidNum = numbers[i]
			log.Println("Part One solution:", invalidNum)
			break
		}
	}

	// Part Two
	set := findContiguousSet(invalidNum, numbers)
	min, max := minAndMax(set)
	log.Println("Part Two solution:", min, "+", max, "=", min+max)
}

func isValid(index int, numbers []int, preambleSize int) bool {
	if index < preambleSize {
		return true
	}
	prevNums := numbers[index-preambleSize : index]
	for i := 0; i < preambleSize; i++ {
		for j := i + 1; j < preambleSize; j++ {
			x, y := prevNums[i], prevNums[j]
			if x+y == numbers[index] {
				return true
			}
		}
	}
	return false
}

func findContiguousSet(invalidNumber int, numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		for j := i + 1; i < len(numbers); j++ {
			sum += numbers[j]
			if sum == invalidNumber {
				return numbers[i : j+1]
			}
			if sum > invalidNumber {
				break
			}
		}
	}
	return []int{}
}

func minAndMax(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		n := numbers[i]
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
