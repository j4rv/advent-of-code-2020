package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part One
	var busIDs []int
	var remainders []int
	splitData := strings.Split(data, "\n")
	startTime, err := strconv.Atoi(splitData[0])
	if err != nil {
		panic("start time not a number")
	}
	for i, rawBusID := range strings.Split(splitData[1], ",") {
		busID, err := strconv.Atoi(rawBusID)
		if err != nil {
			continue
		}
		busIDs = append(busIDs, busID)
		remainders = append(remainders, mod(-i, busID)) // Positive modulo of (-i % busID)
	}
	calcPartOne(startTime, busIDs)

	// Part Two
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
	p := 1 // product of all bus IDs
	for _, busID := range busIDs {
		p *= busID
	}

	s := 0 // ∑(r * (p/a) * modInv(b,r))
	for i := 1; i < len(busIDs); i++ {
		b, r := busIDs[i], rem[i]

		pb := p / b

		inv := modInv(pb, b)

		s += r * pb * inv
	}

	// sanity check
	for i := 0; i < len(busIDs); i++ {
		b, r := busIDs[i], rem[i]
		if s%b != r {
			log.Fatal("CRP implementation broke", s, b, r)
		}
	}

	log.Println("Part Two solution:", s%p)
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

// % is not modulo in golang ლಠ益ಠ)ლ
func mod(val, m int) int {
	val = val % m
	if val < 0 {
		val = m + val
	}
	return val
}
