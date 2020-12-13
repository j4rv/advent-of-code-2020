package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part One
	var busIDs []int
	splitData := strings.Split(data, "\n")
	startTime, err := strconv.Atoi(splitData[0])
	if err != nil {
		panic("start time not a number")
	}
	for _, rawBusID := range strings.Split(splitData[1], ",") {
		busID, err := strconv.Atoi(rawBusID)
		if err != nil {
			continue
		}
		busIDs = append(busIDs, busID)
	}
	calcPartOne(startTime, busIDs)

	// PartTwo
	var remainders []int
	for i, rawBusID := range strings.Split(splitData[1], ",") {
		if rawBusID == "x" {
			continue
		}
		remainders = append(remainders, i)
	}
	calcPartTwo(busIDs, remainders)
}

func calcPartOne(startTime int, busIDs []int) {
	for time := startTime; ; time++ {
		for _, busID := range busIDs {
			if time%busID == 0 {
				waitTime := time - startTime
				log.Println("Part One solution:", waitTime*busID)
				return
			}
		}
	}
}

func calcPartTwo(busIDs []int, rem []int) {
	// Start at bus 0, incrementing by busID[0] at every step
	increment := busIDs[0]
	accum := increment

	for i := 1; i < len(busIDs); i++ {
		for (accum+rem[i])%busIDs[i] != 0 {
			accum += increment
		}
		increment *= busIDs[i] // multiply by current busID, so that accum is always correct for its remainder
	}

	log.Println("Part Two solution:", accum)
}

/*

//Failed attempt to solve it using CRT:

func calcPartTwo(busIDs []int, rem []int) {
	log.Println(busIDs)
	log.Println(rem)

	p := 1 // product of all bus IDs
	for _, busID := range busIDs {
		p *= busID
	}

	s := 0 // ∑(a*pa*c)
	for i := 0; i < len(busIDs); i++ {
		a, b := busIDs[i], rem[i]
		pa := p / a
		c := modInv(pa, a)
		s += b * pa * c
	}
	log.Println(s, p, s%p)
}

// modulo inverse
func modInv(a, m int) int {
	a = a % m
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	panic("mod inv not found")
}
*/
