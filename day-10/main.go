package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	jolts := []int{0} // start with zero (charging outlet)
	for _, dataLine := range strings.Split(data, "\n") {
		num, err := strconv.Atoi(dataLine)
		if err != nil {
			panic("not a number: " + dataLine)
		}
		jolts = append(jolts, num)
	}
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3) // add your device's built-in joltage adapter

	// Part One
	log.Println("Part One solution:", partOne(jolts))

	// Part Two
	log.Println("Part Two solution:", countPossibleAdapters(0, jolts, make(map[int]int)))
}

func partOne(jolts []int) int {
	var diffOne, diffThree int
	for i := range jolts {
		if i == 0 {
			continue
		}
		prev, curr := jolts[i-1], jolts[i]
		if prev+1 == curr {
			diffOne++
		}
		if prev+3 == curr {
			diffThree++
		}
	}
	return diffOne * diffThree
}

func countPossibleAdapters(fromIndex int, nums []int, visited map[int]int) int {
	if fromIndex >= len(nums)-3 {
		return 1
	}

	num := nums[fromIndex]
	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := nums[i]
		if areCompatible(num, n) {
			count += countPossibleAdapters(i, nums, visited)
		}
	}

	visited[num] = count // store the result
	return count
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}
